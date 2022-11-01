package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	qnaming "github.com/saitofun/qkit/x/stringsx"

	g "github.com/saitofun/qkit/gen/codegen"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "strfmt.go", nil, parser.ParseComments)
	file := g.NewFile("strfmt", "strfmt_generated.go")

	regexps := make([]string, 0)
	for key, obj := range f.Scope.Objects {
		if obj.Kind == ast.Con {
			regexps = append(regexps, key)
		}
	}

	snippets := make([]g.Snippet, 0)
	for _, key := range regexps {
		var (
			name          = strings.Replace(key, "regexpString", "", 1)
			validatorName = strings.Replace(qnaming.LowerSnakeCase(name), "_", "-", -1)
			args          = []g.Snippet{g.Ident(key), g.Valuer(validatorName)}
			prefix        = qnaming.UpperCamelCase(name)
			snippet       g.Snippet
		)
		snippet = g.Func().Named("init").Do(
			g.Ref(
				g.Ident(file.Use(pkg, "DefaultFactory")),
				g.Call(
					"Register",
					g.Ident(prefix+"Validator"),
				),
			),
		)
		snippets = append(snippets, snippet)
		snippet = g.DeclVar(
			g.Assign(g.Var(nil, prefix+"Validator")).
				By(g.Call(file.Use(pkg, "NewRegexpStrfmtValidator"), args...)),
		)
		snippets = append(snippets, snippet)

	}
	file.WriteSnippet(snippets...)
	_, _ = file.Write()
}

var pkg = "github.com/saitofun/qkit/kit/validator"
