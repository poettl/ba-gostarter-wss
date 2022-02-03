package wss

import (
	"context"
	"encoding/json"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	upgrader = websocket.Upgrader{}
)

func handleSubscriptions(ctx context.Context, s *api.Server, subscriber *redis.PubSub, ws *websocket.Conn) {
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
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Channel))
			if err != nil {
				errChannel <- err
				return
			}
		}
		// time.Sleep(200 * time.Millisecond)
	}

}

func GetWSSStreamRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1WSS.GET("/stream", getWSSStreamHandler(s))
}

func getWSSStreamHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		ctx := c.Request().Context()
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
			subscriber = s.Redis.Subscribe(ctx, channels...)
			go handleSubscriptions(ctx, s, subscriber, ws)
			// TODO handle error channel
		}

	}
}
