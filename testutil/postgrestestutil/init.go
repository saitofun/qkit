package postgrestestutil

import (
	"net/url"

	"github.com/saitofun/qkit/conf/app"
	"github.com/saitofun/qkit/conf/log"
	"github.com/saitofun/qkit/conf/postgres"
	"github.com/saitofun/qkit/kit/sqlx"
)

type TestEndpoint struct{ postgres.Endpoint }

func (e *TestEndpoint) SetDefault() {
	e.Endpoint.SetDefault()
	if e.Master.Hostname == "" {
		e.Master.Hostname = "127.0.0.1"
	}
	if e.Master.Base == "" {
		e.Master.Base = "test"
	}
	if e.Master.Username == "" {
		e.Master.Username = "test_user"
	}
	if e.Master.Password == "" {
		e.Master.Password = "test_passwd"
	}
	if e.Slave.Hostname == "" {
		e.Slave.Hostname = "localhost"
	}
	if e.Master.Param == nil {
		e.Master.Param = make(url.Values)
	}
	e.Master.Param["sslmode"] = []string{"disable"}
}

var (
	Endpoint = &TestEndpoint{Endpoint: postgres.Endpoint{Database: sqlx.NewDatabase("")}}
)

func init() {
	app.New(
		app.WithName("test"),
		app.WithLogger(log.Std()),
		app.WithRoot("."),
	).Conf(Endpoint)
}
