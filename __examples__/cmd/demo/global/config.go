package global

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/saitofun/qkit/__examples__/pkg/models"
	"github.com/saitofun/qkit/conf/app"
	"github.com/saitofun/qkit/conf/log"
	"github.com/saitofun/qkit/conf/mqtt"
	"github.com/saitofun/qkit/conf/postgres"
)

// global vars

var (
	Postgres = &postgres.Endpoint{Database: models.Demo}
	MQTT     = &mqtt.Broker{}
	Server   = struct{ Port int }{}

	App *app.Ctx
)

func init() {
	App = app.New(
		app.WithName("demo"),
		app.WithRoot(".."),
		app.WithVersion("0.0.1"),
		app.WithLogger(log.Std()),
	)
	App.Conf(Postgres, MQTT, Server)

}

func BatchRun(tasks ...func()) {
	wg := &sync.WaitGroup{}

	for i := range tasks {
		fn := tasks[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}

func StartServer(ctx context.Context) func() {
	return func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello Demo"))
		})
		l := log.FromContext(ctx)

		l.Start(ctx, "http_server")
		defer l.End()

		l.Info("StartServer: http server started at %d", Server.Port)

		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Server.Port), nil)
		if err != nil {
			l.Fatal(err)
		}
	}
}

func EventMonitor(ctx context.Context) func() {
	return func() {
		// TODO monitor chain event
		l := log.FromContext(ctx)
		l.Start(ctx, "event_monitor")
		defer l.End()

		l.Info("EventMonitor: running")
	}
}
