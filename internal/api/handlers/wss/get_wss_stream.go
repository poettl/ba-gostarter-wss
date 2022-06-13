package wss

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/api/auth"
	"allaboutapps.dev/aw/go-starter/internal/models"
	"allaboutapps.dev/aw/go-starter/internal/types/wss"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	upgrader         = websocket.Upgrader{}
	channelMapScopes = map[string]string{"test": auth.AuthScopeApp.String(), "testUser": auth.AuthScopeUser.String()}
)

func handleSubscriptions(ctx context.Context, s *api.Server, subscriber *redis.PubSub, ws *websocket.Conn, userID string) {
	log.Debug().Msg("Subscribing to redis channels")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Error().Msg("Error receiving message")
			return
		}
		log.Debug().Msg("Received message from redis channel \n" + msg.Channel + ": " + msg.Payload)
		select {
		case <-ctx.Done():
			log.Debug().Msg("Context canceled")
			ws.Close()
		default:
			if msg.Channel == userID+"-logout" {
				log.Debug().Msg("User logged out")
				ws.Close()
			} else {
				err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Channel))
				if err != nil {
					log.Error().Msg("Error writing message")
					return
				}
			}

		}
	}

}

func handleForceClose(ctx context.Context, ws *websocket.Conn) {
	time.Sleep(12 * time.Hour)
	log.Debug().Msg("WS force close")
	ws.Close()
}

func GetWSSStreamRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1WSS.GET("/stream", getWSSStreamHandler(s))
}

func getWSSStreamHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := wss.NewGetWSSStreamRouteParams()
		if err := util.BindAndValidateQueryParams(c, &params); err != nil {
			return err
		}

		wssToken, err := models.WSSTokens(
			models.WSSTokenWhere.Token.EQ(string(params.WssToken)),
		).One(ctx, s.DB)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Debug().Msg("No wss token found")
				return echo.NewHTTPError(http.StatusUnauthorized, "No wss token found")
			}
			log.Error().Err(err).Msg("Error getting wss token")
			return err
		}

		if time.Now().After(wssToken.ValidUntil) {
			log.Debug().Msg("Wss token expired")
			return echo.ErrUnauthorized
		}

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		_, err = wssToken.Delete(ctx, s.DB)
		if err != nil {
			log.Error().Err(err).Msg("Error deleting wss token")
			return err
		}

		user, err := models.FindUser(ctx, s.DB, wssToken.UserID)
		if err != nil {
			log.Error().Err(err).Msg("Error finding user")
			return err
		}

		var subscriber *redis.PubSub
		defer ws.Close()
		go handleForceClose(ctx, ws)
		for {
			// Read
			_, msg, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}
			log.Debug().Msg("Received message from client: " + string(msg))

			// handle subscriptions
			if subscriber != nil {
				log.Debug().Msg("Subscription to redis channel already exists")
				subscriber.Close()
				subscriber = nil
			}
			channels := []string{}
			if err := json.Unmarshal(msg, &channels); err != nil {
				log.Err(err).Msg("Failed to unmarshal channels")
				ws.Close()
				return err
			}

			for _, item := range channels {
				scope, ok := channelMapScopes[item]
				if !ok {
					log.Debug().Msg("Channel not found")
					ws.Close()
					return echo.ErrUnauthorized
				}
				if !util.ContainsString(user.Scopes, scope) {
					log.Debug().Msg("User not authorized")
					ws.Close()
					return echo.ErrUnauthorized
				}

			}
			channels = append(channels, wssToken.UserID+"-logout")
			subscriber = s.Redis.Subscribe(ctx, channels...)
			go handleSubscriptions(ctx, s, subscriber, ws, wssToken.UserID)
		}

	}
}
