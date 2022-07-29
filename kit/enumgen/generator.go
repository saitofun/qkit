package enumgen

import (
	"go/types"
	"log"
	"path"
	"path/filepath"
	"sort"

	"github.com/saitofun/qkit/gen/codegen"
	"github.com/saitofun/qkit/x/pkgx"
	"github.com/saitofun/qlib/util/qnaming"
	"golang.org/x/tools/go/packages"
)

type Generator struct {
	pkg *pkgx.Pkg
	*Scanner
	enums map[*types.TypeName]*Enum
}

func New(pkg *pkgx.Pkg) *Generator {
	return &Generator{
		pkg:     pkg,
		Scanner: NewScanner(pkg),
		enums:   map[*types.TypeName]*Enum{},
	}
}

func (g *Generator) Scan(names ...string) {
	for _, name := range names {
		tn := g.pkg.TypeName(name)
		if tn == nil {
			continue
		}
		opts, ok := g.Scanner.Options(tn)
		if ok && opts[0].Int != nil && opts[0].Str != nil {
			sort.Slice(opts, func(i, j int) bool {
				return *opts[i].Int < *opts[j].Int
			})
			g.enums[tn] = NewEnum(tn.Pkg().Path()+"."+tn.Name(), opts)
		}
	}
}

func (g Generator) Output(cwd string) {
	for tn, enum := range g.enums {
		dir, _ := filepath.Rel(cwd, pkgDir(tn.Pkg().Path()))
		filename := codegen.GenerateFileSuffix(
			path.Join(dir, qnaming.LowerSnakeCase(enum.Name)+".go"))
		f := codegen.NewFile(tn.Pkg().Name(), filename)
		enum.WriteToFile(f)

		if _, err := f.Write(); err != nil {
			log.Printf("%s generate failed: %v", filename, err)
		}
	}
}

// _enum test only
func _enum(g *Generator, name string) *Enum {
	for tn, enum := range g.enums {
		if tn.Name() == name {
			return enum
		}
	}
	return nil
}

func pkgDir(path string) string {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles,
	}, path)
	if err != nil {
		panic(err)
	}
	if len(pkgs) == 0 {
		panic("package `" + path + "` not found")
	}
	return filepath.Dir(pkgs[0].GoFiles[0])
}
