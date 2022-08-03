package http

import (
	"context"
	"net/http"
)

func NewContextInjectorMw(ctx WithContext) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				if ctx != nil {
					req = req.WithContext(ctx(req.Context()))
				}
				next.ServeHTTP(rw, req)
			},
		)
	}
}

type (
	WithContext = func(ctx context.Context) context.Context
	Middleware  = func(http.Handler) http.Handler
)

func WithContextCompose(contexts ...WithContext) WithContext {
	return func(ctx context.Context) context.Context {
		for i := range contexts {
			ctx = contexts[i](ctx)
		}
		return ctx
	}
}
