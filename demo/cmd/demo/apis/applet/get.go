package applet

import (
	"context"

	"github.com/saitofun/qkit/demo/pkg/modules/applet"
	"github.com/saitofun/qkit/kit/httptransport/httpx"
)

type ListApplet struct {
	httpx.MethodGet `summary:"get applet info"`
	applet.ListAppletReq
}

func (r *ListApplet) Output(ctx context.Context) (interface{}, error) {
	return applet.ListApplets(ctx, &r.ListAppletReq)
}
