package impl

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"

	"strapi-webhook/base"
)

func Middlewares(config base.ServerConfig, logger log.Logger) map[string]endpoint.Middleware {
	middlewares := map[string]endpoint.Middleware{
		"entry": entryMiddleware(config, logger),
		"media": mediaMiddleware(config, logger),
	}
	return middlewares
}

func loggingMiddleware(name string, logger log.Logger) endpoint.Middleware {
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
func entryMiddleware(config base.ServerConfig, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (resp any, err error) {
			endpoint := loggingMiddleware("Entry", logger)(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
func mediaMiddleware(config base.ServerConfig, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (resp any, err error) {
			endpoint := loggingMiddleware("Media", logger)(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
