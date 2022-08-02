package cmd

import (
	"github.com/spf13/cobra"

	"github.com/saitofun/qkit/kit/statusxgen"
	"github.com/saitofun/qkit/x/pkgx"
)

func init() {
	cmd := &cobra.Command{
		Use:     "status",
		Aliases: []string{"status-error", "error"},
		Short:   "generate interfaces of status error",
		Run: func(cmd *cobra.Command, args []string) {
			run("status", func(pkg *pkgx.Pkg) Generator {
				g := statusxgen.New(pkg)
				g.Scan(args...)
				return g
			}, args...)
		},
	}

	Cmd.AddCommand(cmd)
}
