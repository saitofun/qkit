package main

import (
	"github.com/saitofun/qkit/httptransporttestutil/testdata/server/cmd/app/routes"
	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/kit"
)

func main() {
	ht := &httptransport.HttpTransport{
		Port: 8080,
	}
	ht.SetDefault()

	kit.Run(routes.RootRouter, ht)
}
