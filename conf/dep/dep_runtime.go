package dep

import (
	"github.com/sincospro/qkit/base/cmd"
	"github.com/sincospro/qkit/conf/dep/dep_enums"
)

type Runtime struct {
	Type      dep_enums.RuntimeType   `json:"type"`
	Method    dep_enums.RuntimeMethod `json:"method"`
	Config    string                  `json:"config"`
	StartCmd  cmd.Command             `json:"start_cmd"`
	StopCmd   cmd.Command             `json:"stop_cmd"`
	StatusCmd cmd.Command             `json:"status_cmd"`
	Envs      [][2]string             `json:"envs"`
	LogRoot   string                  `json:"log_root"`
}

func (r *Runtime) Start() error  { return nil }
func (r *Runtime) Stop() error   { return nil }
func (r *Runtime) Status() error { return nil }
