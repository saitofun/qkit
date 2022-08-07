package main

import (
	"sync"
	"time"

	"github.com/saitofun/qkit/demo/cmd/demo/apis"
	"github.com/saitofun/qkit/demo/cmd/demo/global"
	"github.com/saitofun/qkit/kit/kit"
)

var app = global.App

func main() {
	app.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})

	// TODO should add batch routines/daemons to app context
	app.Execute(func(args ...string) {
		BatchRun(
			func() {
				kit.Run(apis.RouterRoot, global.Server())
			},
		)
	})
}

func BatchRun(commands ...func()) {
	wg := &sync.WaitGroup{}

	for i := range commands {
		cmd := commands[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			cmd()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
