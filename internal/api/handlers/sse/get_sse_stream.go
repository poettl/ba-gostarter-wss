package sse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/types/sse"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/labstack/echo/v4"
)

func GetSSEStreamRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1SSE.GET("/stream", getSSEStreamHandler(s))
}

func getSSEStreamHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := sse.NewGetSSEStreamRouteParams()
		if err := util.BindAndValidateQueryParams(c, &params); err != nil {
			return err
		}
		// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		c.Response().WriteHeader(http.StatusOK)
		if len(params.ChannelList) == 0 {
			fmt.Println("No channels provided")
			return c.NoContent(http.StatusNoContent)
		}
		subscriber := s.Redis.Subscribe(ctx, params.ChannelList...)
		for {
			msg, err := subscriber.ReceiveMessage(ctx)
			if err != nil {
				return err
			}

			fmt.Println("Received message from " + msg.Channel + " channel.")
			fmt.Printf("%+v\n", msg.Payload)

			// timeout := time.After(1 * time.Second)
			select {
			case <-ctx.Done():
				log.Println("Context canceled")
				return nil
			// case <-timeout:
			// 	if _, err := fmt.Fprintf(c.Response().Writer, ": nothing to sent\n\n"); err != nil {
			// 		return err
			// 	}
			default:
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				if err := enc.Encode(msg.Channel); err != nil {
					return err
				}

				if _, err := fmt.Fprintf(c.Response().Writer, "data: %v\n\n", buf.String()); err != nil {
					return err
				}
				fmt.Printf("data: %v\n", buf.String())
				// if _, err := c.Response().Write(buf.Bytes()); err != nil {
				// 	return err
				// }
				c.Response().Flush()
			}
		}

	}

}
