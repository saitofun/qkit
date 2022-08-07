package applet

import (
	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/kit"
)

var Root = kit.NewRouter(httptransport.Group("/applet"))

func init() {
	Root.Register(kit.NewRouter(&CreateApplet{}))
	Root.Register(kit.NewRouter(&ListApplet{}))
}
