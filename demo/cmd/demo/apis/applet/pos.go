package applet

import (
	"context"

	"github.com/saitofun/qkit/demo/pkg/modules/applet"
	"github.com/saitofun/qkit/kit/httptransport/httpx"
)

type CreateApplet struct {
	httpx.MethodPost             `summary:"create applet by name"`
	applet.CreateAppletByNameReq `in:"body"`
}

func (r *CreateApplet) Output(ctx context.Context) (interface{}, error) {
	return applet.CreateAppletByName(ctx, &r.CreateAppletByNameReq)
}
