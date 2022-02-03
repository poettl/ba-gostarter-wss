package wss

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func handleSubscriptions(ctx context.Context, s *api.Server, msg []byte, subscriber *redis.PubSub, ws *websocket.Conn, errChannel chan error) {
	if subscriber != nil {
		subscriber.Close()
	}
	fmt.Println("starting subscription goroutine")
	channels := []string{}
	if err := json.Unmarshal(msg, &channels); err != nil {
		errChannel <- err
		ws.Close()
		return
	}
	if len(channels) == 0 {
		return
	}
	subscriber = s.Redis.Subscribe(ctx, channels...)
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			errChannel <- err
			return
		}
		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", msg.Payload)
		select {
		case <-ctx.Done():
			log.Println("Context canceled")
			ws.Close()
		default:
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Channel))
			if err != nil {
				errChannel <- err
				return
			}
		}
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
		errChannel := make(chan error)
		var subscriber *redis.PubSub
		err = <-errChannel
		if err != nil {
			return err
		}
		defer ws.Close()
		var msg []byte
		for {
			go handleSubscriptions(ctx, s, msg, subscriber, ws, errChannel)
			// Read
			_, msg, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}
}
