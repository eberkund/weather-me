module github.com/eberkund/weather-me

go 1.23.2

require (
	connectrpc.com/connect v1.17.0
	connectrpc.com/cors v0.1.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/AlekSi/pointer v1.2.0
	github.com/antihax/optional v1.0.0
	github.com/eberkund/graceful v0.2.0
	github.com/go-faster/errors v0.7.1
	github.com/justinas/alice v1.2.0
	github.com/rs/cors v1.11.1
	github.com/rs/zerolog v1.33.0
	github.com/sethvargo/go-envconfig v1.1.0
	golang.org/x/net v0.23.0
	golang.org/x/oauth2 v0.23.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/xid v1.5.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/eberkund/weather-me => ./
