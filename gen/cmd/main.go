package main

import (
	"os"

	"github.com/saitofun/qkit/gen/cmd/gen"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "gen",
	Version: "0.0.1",
}

func init() {
	verbose := false
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "")
	cmd.AddCommand(gen.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		cmd.PrintErr(err)
		os.Exit(-1)
	}
}
