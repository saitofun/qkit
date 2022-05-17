package deps_test

import (
	"testing"

	"github.com/sincospro/qkit/base/types"
	"github.com/sincospro/qkit/dep"
)

type PostgresqlConfig struct {
	Command               string         `name:"command,default=''"`
	User                  string         `name:"user,default='postgres'"`
	AutoStart             bool           `name:"autostart,default=false"`
	AutoRestart           bool           `name:"atuorestart,default=true"`
	StartSecs             types.Second   `name:"startsecs"`
	StartRetries          int            `name:"startretries"`
	ExitCodes             int            `name:"exitcodes"`
	Priority              int            `name:"priority"`
	StopSignal            types.Signal   `name:"stopsignal"`
	Directory             string         `name:"directory"`
	RedirectStderr        bool           `name:"redirectstderr"`
	StdoutLogFile         string         `name:"stdout_logfile"`
	StdoutLogFileMaxBytes types.Capacity `name:"stdout_logfile_maxbytes"`
	StdoutLogFileBackups  int            `name:"stdout_logfile_backups"`
	StderrLogFile         string         `name:"stderr_logfile"`
}

func TestConfigLoader_Load(t *testing.T) {
	loader := &deps.ConfigLoader{
		Path: "./__tests__/conf/postgresql.conf",
		Sep:  '=',
	}
	err := loader.Load()
	if err != nil {
		t.Log(err)
	}

	for k, v := range loader.Values {
		t.Logf("%s=%s", k, v)
	}
}
