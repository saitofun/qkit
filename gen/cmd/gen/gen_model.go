package gen

import (
	"github.com/go-courier/packagesx"
	"github.com/go-courier/sqlx/v2/generator"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "generate database model",
		Run: func(cmd *cobra.Command, args []string) {
			if database == "" {
				panic("database required")
			}
			for _, arg := range args {
				Run("model", func(pkg *packagesx.Package) Generator {
					g := generator.NewSqlFuncGenerator(pkg)
					g.WithComments = true
					g.WithTableInterfaces = true
					g.StructName = arg
					g.Database = database
					g.TableName = tableName

					g.Scan()
					return g
				})
			}
		},
	}
	cmd.Flags().StringVarP(
		&database, "database", "", "", "(required) database name",
	)
	cmd.Flags().StringVarP(
		&tableName, "table-name", "t", "", "custom table name",
	)
	cmd.Flags().BoolVarP(
		&withMethods, "with-methods", "", true, "with table methods",
	)

	Cmd.AddCommand(cmd)
}

var (
	database    string
	tableName   string
	withMethods bool
)
