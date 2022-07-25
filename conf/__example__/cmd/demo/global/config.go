package global

import (
	"fmt"

	"github.com/saitofun/qkit/conf/__example__/pkg/models"
	"github.com/saitofun/qkit/conf/app"
	"github.com/saitofun/qkit/conf/mqtt"
	"github.com/saitofun/qkit/conf/postgres"
	"github.com/saitofun/qlib/encoding/qjson"
)

// global vars

var (
	Postgres = &postgres.Endpoint{Database: models.Demo}
	MQTT     = &mqtt.Broker{}
)

func init() {
	app := app.New(
		app.WithName("demo"),
		app.WithRoot(".."),
	)
	app.Conf(Postgres, MQTT)

	fmt.Println(qjson.UnsafeMarshalIndent(Postgres))
	fmt.Println(qjson.UnsafeMarshalIndent(MQTT))
}
