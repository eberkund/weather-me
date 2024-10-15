package setup

import (
	"context"
	"net/http"

	"github.com/eberkund/graceful"
	"github.com/eberkund/weather-me/pkg/server"
)

type App struct {
	*graceful.Graceful
}

func New(builder *server.Builder) *App {
	g := graceful.New(context.Background())
	var s *http.Server
	g.Go(func(ctx context.Context) error {
		s = builder.Build(ctx)
		return s.ListenAndServe()
	})
	g.Stop(func(ctx context.Context) error {
		return s.Shutdown(ctx)
	})
	return &App{
		Graceful: g,
	}
}
