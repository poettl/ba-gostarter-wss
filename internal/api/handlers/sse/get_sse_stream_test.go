package sse_test

import (
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/test"
)

func TestGetSSEStream(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {

		test.PerformRequest(t, s, "GET", "/api/v1/sse/stream", nil, nil)

		// assert.Equal(t, http.StatusNoContent, res.Result().StatusCode)
	})
}
