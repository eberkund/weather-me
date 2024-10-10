package server

import (
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// LogRequests returns middleware that logs every HTTP request.
func LogRequests() func(next http.Handler) http.Handler {
	c := alice.New(
		hlog.NewHandler(log.Logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status_code", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("incoming request")
		}),
	)
	return c.Then
}
