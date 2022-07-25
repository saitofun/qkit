package cmd

import (
	"github.com/saitofun/qkit/kit/enumgen"
	"github.com/saitofun/qkit/x/pkgx"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "enum",
		Short: "generate interfaces of enumeration",
		Run: func(cmd *cobra.Command, args []string) {
			run("enum", func(pkg *pkgx.Pkg) Generator {
				g := enumgen.New(pkg)
				g.Scan(args...)
				return g
			})
		},
	})
}
