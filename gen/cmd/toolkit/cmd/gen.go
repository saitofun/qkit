package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/saitofun/qkit/x/misc/timer"
	"github.com/saitofun/qkit/x/pkgx"
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
		fmt.Println(cwd)
		r := color.New(color.FgHiRed).SprintfFunc()
		g := color.New(color.FgHiGreen).SprintfFunc()
		b := color.New(color.FgHiBlue).SprintfFunc()
		fmt.Printf("%s : %s %s: cost %s",
			r(cmd),
			g(pkg.Name),
			b("%6dms", cost().Milliseconds()),
		)
	}()

	g(pkg).Output(cwd)
}
