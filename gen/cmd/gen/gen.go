package gen

import (
	"log"
	"os"

	"github.com/saitofun/qkit/x/misc/timer"
	"github.com/saitofun/qkit/x/pkgx"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "generators",
}

type Generator interface {
	Output(cwd string)
}

func run(cmd string, g func(*pkgx.Pkg) Generator) {
	cwd, _ := os.Getwd()
	cost := timer.Start()

	pkg, err := pkgx.LoadFrom(cwd)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Printf("%s %s: cost %s", cmd, pkg.ID, cost().String())
	}()

	g(pkg).Output(cwd)
}
