package sse

import (
	"net/http"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/types/sse"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/labstack/echo/v4"
)

func GetSSETriggerRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1SSE.GET("/trigger/:value", getSSETriggerHandler(s))
}

func getSSETriggerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := sse.NewGetSSETriggerRouteParams()
		if err := util.BindAndValidatePathParams(c, &params); err != nil {
			return err
		}

		if err := s.Redis.Publish(ctx, "test", params.Value).Err(); err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}
