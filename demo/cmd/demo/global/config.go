package global

import (
	"context"

	confapp "github.com/saitofun/qkit/conf/app"
	confhttp "github.com/saitofun/qkit/conf/http"
	confjwt "github.com/saitofun/qkit/conf/jwt"
	conflog "github.com/saitofun/qkit/conf/log"
	confmqtt "github.com/saitofun/qkit/conf/mqtt"
	confpostgres "github.com/saitofun/qkit/conf/postgres"
	"github.com/saitofun/qkit/demo/pkg/models"
	"github.com/saitofun/qkit/kit/kit"
	"github.com/saitofun/qkit/kit/sqlx/migration"
	"github.com/saitofun/qkit/x/contextx"
)

// global vars

var (
	postgres = &confpostgres.Endpoint{Database: models.Demo}
	mqtt     = &confmqtt.Broker{}
	server   = &confhttp.Server{}
	jwt      = &confjwt.Jwt{}
	logger   = &conflog.Log{Name: "srv-demo"}
	std      = conflog.Std()

	App *confapp.Ctx
)

func init() {
	App = confapp.New(
		confapp.WithName("demo"),
		confapp.WithRoot(".."),
		confapp.WithVersion("0.0.1"),
		confapp.WithLogger(conflog.Std()),
	)
	App.Conf(postgres, server, jwt, logger)

	confhttp.RegisterCheckerBy(postgres, mqtt, server)
	std.(conflog.LevelSetter).SetLevel(conflog.InfoLevel)
}

var WithContext = contextx.WithContextCompose(
	WithDatabase(postgres),
	WithLogger(conflog.Std()),
)

func Server() kit.Transport {
	return server.WithContextInjector(WithContext)
}

func Migrate() {
	ctx, log := conflog.StdContext(context.Background())

	log.Start(ctx, "Migrate")
	defer log.End()
	if err := migration.Migrate(postgres.WithContext(ctx), nil); err != nil {
		log.Panic(err)
	}
}
