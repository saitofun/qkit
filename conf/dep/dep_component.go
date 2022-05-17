package dep

import (
	"context"

	"github.com/sincospro/qkit/base/cmd"
	"github.com/sincospro/qkit/base/ver"
	"github.com/sincospro/qkit/conf/dep/dep_enums"
)

type Component struct {
	Name         dep_enums.ComponentName `json:"name"`
	Type         dep_enums.ComponentType `json:"type"`
	Version      ver.Version             `json:"version"`
	InstallCmd   cmd.Command             `json:"install_cmd"`
	UninstallCmd cmd.Command             `json:"uninstall_cmd"`
	Runtime      Runtime                 `json:"runtime"`
}

func (c *Component) Install(ctx context.Context) {}

func (c *Component) Uninstall(ctx context.Context) {}

func (c *Component) Daemonize(ctx context.Context) {}
