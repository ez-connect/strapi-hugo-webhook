package impl

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"strapiwebhook/base"
)

func Middlewares(config base.ServerConfig) map[string]endpoint.Middleware {
	middlewares := map[string]endpoint.Middleware{
		"entry": entryMiddleware(config),
		"media": mediaMiddleware(config),
	}
	return middlewares
}

func loggingMiddleware(name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (resp any, err error) {
			resp, err = next(ctx, request)
			if err != nil {
				// nolint:errcheck
				logger.Log("endpoint", name, "err", err)
			}
			return
		}
	}
}
func entryMiddleware(config base.ServerConfig) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (resp any, err error) {
			endpoint := loggingMiddleware("Entry")(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
func mediaMiddleware(config base.ServerConfig) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (resp any, err error) {
			endpoint := loggingMiddleware("Media")(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
