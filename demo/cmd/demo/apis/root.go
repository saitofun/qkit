package apis

import (
	"github.com/saitofun/qkit/conf/http"
	"github.com/saitofun/qkit/demo/cmd/demo/apis/applet"
	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/kit"
)

var (
	RouterRoot = kit.NewRouter(httptransport.Group("/"))
	RouterV0   = kit.NewRouter(httptransport.BasePath("/w3bstream/v0"))
)

func init() {
	RouterRoot.Register(http.LivenessRouter)
	RouterRoot.Register(RouterV0)

	RouterV0.Register(applet.Root)
}
