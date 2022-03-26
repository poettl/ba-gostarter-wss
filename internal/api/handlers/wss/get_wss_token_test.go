package wss_test

import (
	"net/http"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/test"
	"allaboutapps.dev/aw/go-starter/internal/types"
	"github.com/stretchr/testify/require"
)

func TestGetWSSToken_Success(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/wss/token", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetWSSTokenResponse
		test.ParseResponseAndValidate(t, res, &response)

		require.NotNil(t, response.Token)
	})
}
