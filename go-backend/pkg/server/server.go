package server

import (
	"context"
	"net"
	"net/http"

	connectcors "connectrpc.com/cors"
	"connectrpc.com/grpcreflect"
	"github.com/eberkund/weather-me/gen/weather/v1/weatherv1connect"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Builder struct {
	weather *WeatherHandler
	cors    *cors.Cors
	port    int
}

func NewServer(weather *WeatherHandler, port int, origins []string) *Builder {
	return &Builder{
		weather: weather,
		port:    port,
		cors: cors.New(cors.Options{
			AllowedOrigins: origins,
			AllowedMethods: connectcors.AllowedMethods(),
			AllowedHeaders: append(connectcors.AllowedHeaders(), "Authorization"),
			ExposedHeaders: connectcors.ExposedHeaders(),
		}),
	}
}

func (n *Builder) Build(ctx context.Context) *http.Server {
	mux := http.NewServeMux()

	mux.Handle(weatherv1connect.NewWeatherServiceHandler(n.weather))

	reflector := grpcreflect.NewStaticReflector(weatherv1connect.WeatherServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	chain := alice.New(
		LogRequests(),
		n.cors.Handler,
	).Then(mux)
	addr := &net.TCPAddr{Port: n.port}
	return &http.Server{
		Addr:    addr.String(),
		Handler: h2c.NewHandler(chain, &http2.Server{}),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
}
