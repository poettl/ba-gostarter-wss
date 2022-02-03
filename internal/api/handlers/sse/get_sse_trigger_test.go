package sse_test

import (
	"net/http"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSEETrigger(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {

		res := test.PerformRequest(t, s, "GET", "/api/v1/sse/trigger/test", nil, nil)
		assert.Equal(t, http.StatusNoContent, res.Result().StatusCode)
	})
}
