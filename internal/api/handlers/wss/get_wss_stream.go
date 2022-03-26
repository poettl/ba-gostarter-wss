package wss

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/models"
	"allaboutapps.dev/aw/go-starter/internal/types/wss"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	upgrader = websocket.Upgrader{}
)

func handleSubscriptions(ctx context.Context, s *api.Server, subscriber *redis.PubSub, ws *websocket.Conn, userID string) {
	errChannel := make(chan error)
	log.Debug().Msg("Subscribing to redis channels")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			errChannel <- err
			return
		}
		log.Debug().Msg("Received message from redis channel \n" + msg.Channel + ": " + msg.Payload)
		select {
		case <-ctx.Done():
			log.Debug().Msg("Context canceled")
			ws.Close()
		case <-errChannel:
			log.Error().Msg("Error receiving message")
		default:

			if msg.Channel == userID+"-logout" {
				log.Debug().Msg("User logged out")
				ws.Close()
			} else {
				err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Channel))
				if err != nil {
					errChannel <- err
					return
				}
			}

		}
	}

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

		var subscriber *redis.PubSub
		defer ws.Close()
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
			channels = append(channels, wssToken.UserID+"-logout")
			subscriber = s.Redis.Subscribe(ctx, channels...)
			go handleSubscriptions(ctx, s, subscriber, ws, wssToken.UserID)
		}

	}
}
