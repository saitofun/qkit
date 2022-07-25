package gen

import (
	"log"
	"os"
	"time"

	"github.com/go-courier/packagesx"
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

	ts := time.Now()
	pkg, err := pkgx.LoadFrom(cwd)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Printf("%s %s: cost %s", cmd, pkg.ID, time.Since(ts))
	}()

	g(pkg).Output(cwd)
}

func Run(cmd string, g func(*packagesx.Package) Generator) {
	cwd, _ := os.Getwd()

	ts := time.Now()
	pkg, err := packagesx.Load(cwd)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Printf("%s %s: cost %s", cmd, pkg.ID, time.Since(ts))
	}()

	g(pkg).Output(cwd)
}
