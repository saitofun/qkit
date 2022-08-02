package main

import (
	"github.com/saitofun/qkit/kit/httptransport"
	"github.com/saitofun/qkit/kit/kit"
	"github.com/saitofun/qkit/testutil/httptransporttestutil/server/cmd/app/routes"
)

func main() {
	ht := &httptransport.HttpTransport{
		Port: 8080,
	}
	ht.SetDefault()

	kit.Run(routes.RootRouter, ht)
}
