package sse

import (
	"encoding/json"
	"log"
	"net/http"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"github.com/labstack/echo/v4"
)

func GetSEEStreamRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1SSE.GET("/stream", getSEEStreamHandler(s))
}

func getSEEStreamHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		subscriber := s.Redis.Subscribe(ctx, "test")

		enc := json.NewEncoder(c.Response())

		msg, err := subscriber.ReceiveMessage(ctx)

		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				if err != nil {
					return err
				}

				log.Printf("new msg: %+v\n", msg.Payload)
				if err := enc.Encode(msg); err != nil {
					return err
				}
				c.Response().Flush()
			}

		}

	}
}
