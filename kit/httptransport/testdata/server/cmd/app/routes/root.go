package routes

import (
	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/httptransport/swagger"
	"github.com/saitofun/qkit/kit/kit"
)

var RootRouter = kit.NewRouter(httptransport.BasePath("/demo"))

func init() {
	RootRouter.Register(swagger.Router)
}
