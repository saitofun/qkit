package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/kit/kit"
)

var CookieRouter = kit.NewRouter(httptransport.Group("/cookie"))

func init() {
	RootRouter.Register(CookieRouter)

	CookieRouter.Register(kit.NewRouter(&Cookie{}))
}

type Cookie struct {
	httpx.MethodPost
	Token string `name:"token,omitempty" in:"cookie"`
}

func (req *Cookie) Output(ctx context.Context) (interface{}, error) {
	return httpx.Compose(
		httpx.WrapCookies(&http.Cookie{
			Name:    "token",
			Value:   req.Token,
			Expires: time.Now().Add(24 * time.Hour),
		}),
		httpx.WrapStatusCode(http.StatusNoContent),
	)(nil), nil
}
