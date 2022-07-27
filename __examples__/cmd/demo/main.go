package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/saitofun/qkit/__examples__/cmd/demo/global"
	"github.com/spf13/cobra"
)

var app = global.App

func main() {
	app.AddCommand("config", func(...string) {
		spew.Dump(global.Postgres)
		spew.Dump(global.MQTT)
		spew.Dump(global.Server)
	}, func(cmd *cobra.Command) {
		cmd.Short = "dump app config value"
	})

	app.AddCommand("version", func(...string) {
		fmt.Println(app.String())
	}, func(cmd *cobra.Command) {
		cmd.Short = "print app version"
	})

	app.AddCommand("gen", func(...string) {
		if err := app.MarshalDefault(); err != nil {
			panic(err)
		}
	}, func(cmd *cobra.Command) {
		cmd.Short = "generate app default config"
	})

	// TODO gen deployer config
	app.AddCommand("docker", func(...string) {
		fmt.Println("todo: generate app docker config")
	}, func(cmd *cobra.Command) {
		cmd.Short = "generate app docker config"
	})

	// app.AddCommand("supervisor", func(...string) {
	// 	program.SetDefault()
	// 	filename := path.Join(app.Root(), "supervisor.conf")
	// 	if err := program.Write(filename); err != nil {
	// 		panic(fmt.Errorf("init `supervisor.conf` error: %v", err))
	// 	}
	// 	data, _ := program.Bytes()
	// 	fmt.Println(string(data))
	// }, func(cmd *cobra.Command) {
	// 	cmd.Short = "generate app supervisor config"
	// })

	// TODO should add batch routines/daemons to app context

	app.AddCommand("run", func(...string) {
		global.BatchRun(
			global.StartServer(app.Context()),
			global.EventMonitor(app.Context()),
		)
	})
}
