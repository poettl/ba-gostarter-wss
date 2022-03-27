package wss

import (
	"net/http"
	"time"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/api/auth"
	"allaboutapps.dev/aw/go-starter/internal/api/middleware"
	"allaboutapps.dev/aw/go-starter/internal/models"
	"allaboutapps.dev/aw/go-starter/internal/types"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"allaboutapps.dev/aw/go-starter/internal/util/db"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetWSSTokenRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1WSS.GET("/token", getWSSTokenHandler(s))
}

func getWSSTokenHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user := auth.UserFromContext(ctx)

		if !user.IsActive {
			log.Debug().Msg("User is deactivated, rejecting authentication")
			return middleware.ErrForbiddenUserDeactivated
		}

		response := &types.GetWSSTokenResponse{}

		if err := db.WithTransaction(ctx, s.DB, func(tx boil.ContextExecutor) error {
			wssToken := models.WSSToken{
				ValidUntil: time.Now().Add(s.Config.Auth.WSSTokenValidity),
				UserID:     user.ID,
			}

			if err := wssToken.Insert(ctx, tx, boil.Infer()); err != nil {
				log.Debug().Err(err).Msg("Failed to insert wss token")
				return err
			}
			response.Token = (*strfmt.UUID)(&wssToken.Token)

			return nil
		}); err != nil {
			return err
		}
		return util.ValidateAndReturn(c, http.StatusOK, response)
	}
}
