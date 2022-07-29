package httptransport

import (
	"context"
	"net/http"

	"github.com/saitofun/qkit/kit/kit"
	"github.com/saitofun/qkit/x/contextx"
)

type ServiceMeta struct {
	Name    string
	Version string
}

type (
	ckHttpRequest     struct{}
	ckServiceMeta     struct{}
	ckOperatorID      struct{}
	ckOperatorFactory struct{}
)

func ContextWithHttpRequest(ctx context.Context, req *http.Request) context.Context {
	return contextx.WithValue(ctx, ckHttpRequest{}, req)
}

func HttpRequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(ckHttpRequest{}).(*http.Request)
	return r
}

func ContextWithServiceMeta(ctx context.Context, meta *ServiceMeta) context.Context {
	return contextx.WithValue(ctx, ckServiceMeta{}, meta)
}

func ServiceMetaFromContext(ctx context.Context) *ServiceMeta {
	r, _ := ctx.Value(ckServiceMeta{}).(*ServiceMeta)
	return r
}

func ContextWithOperatorID(ctx context.Context, id string) context.Context {
	return contextx.WithValue(ctx, ckOperatorID{}, id)
}

func OperatorIDFromContext(ctx context.Context) string {
	r, _ := ctx.Value(ckOperatorID{}).(string)
	return r
}

func ContextWithOperatorFactory(ctx context.Context, factory *kit.OperatorFactory) context.Context {
	return contextx.WithValue(ctx, ckOperatorFactory{}, factory)
}

func OperatorFactoryFromContext(ctx context.Context) *kit.OperatorFactory {
	r, _ := ctx.Value(ckOperatorFactory{}).(*kit.OperatorFactory)
	return r
}
