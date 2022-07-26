package cmd

import (
	"github.com/saitofun/qkit/kit/statusxgen"
	"github.com/saitofun/qkit/x/pkgx"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "generate interfaces of status error",
		Run: func(cmd *cobra.Command, args []string) {
			run("status", func(pkg *pkgx.Pkg) Generator {
				g := statusxgen.New(pkg)
				g.Scan(args...)
				return g
			})
		},
	}

	Cmd.AddCommand(cmd)
}
