package impl

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Middleware = map[string]func(logger log.Logger) endpoint.Middleware

var Middlewares Middleware = Middleware{
	"Entry": EntryMiddleware,
	"Media": MediaMiddleware,
}

func loggingMiddleware(name string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			resp, err = next(ctx, request)
			if err != nil {
				// nolint:errcheck
				logger.Log("endpoint", name, "err", err)
			}
			return
		}
	}
}
func EntryMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			// defer GetMetrics().EntryMetricFunction(time.Now())
			endpoint := loggingMiddleware("Entry", logger)(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
func MediaMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			// defer GetMetrics().MediaMetricFunction(time.Now())
			endpoint := loggingMiddleware("Media", logger)(next)
			resp, err = endpoint(ctx, request)
			return
		}
	}
}
