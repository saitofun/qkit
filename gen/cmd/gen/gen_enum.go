package gen

import (
	"github.com/saitofun/qkit/infra/enum/generator"
	"github.com/saitofun/qkit/x/pkgx"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "enum",
		Short: "generate interfaces of enumeration",
		Run: func(cmd *cobra.Command, args []string) {
			run("enum", func(pkg *pkgx.Pkg) Generator {
				g := generator.New(pkg)
				g.Scan(args...)
				return g
			})
		},
	})
}
