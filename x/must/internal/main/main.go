package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	. "github.com/saitofun/qkit/gen/codegen"
	"github.com/saitofun/qlib/util/qnaming"
)

func main() {
	pkg := "must"
	_, path, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(path), "../..")

	if name := filepath.Base(root); name != pkg {
		log.Panicf("wrong execute location: \n\tpath: %s\n\tbase: %s", root, name)
	}

	{
		filename := filepath.Join(root, "must.go")
		_, err := os.Stat(filename)
		if !os.IsExist(err) {
			file := NewFile(pkg, filename)
			file.WriteSnippet(
				Comments("pls add your assert function here, or add type and re-generate"),
			)
			if _, err := file.Write(); err != nil {
				log.Panic(err)
			}
		}
	}

	types := []struct {
		BuiltInType
		Name string
	}{
		{BuiltInType: "[]byte", Name: "bytes"},
		{BuiltInType: "string"},
		{BuiltInType: "[]string", Name: "strings"},
		{BuiltInType: "int"},
		{BuiltInType: "int8"},
		{BuiltInType: "int16"},
		{BuiltInType: "int32"},
		{BuiltInType: "int64"},
		{BuiltInType: "uint8"},
		{BuiltInType: "uint16"},
		{BuiltInType: "uint32"},
		{BuiltInType: "uint64"},
		{BuiltInType: "rune"},
		{BuiltInType: "float32"},
		{BuiltInType: "float64"},
	}

	{
		filename := filepath.Join(root, "must_generated.go")
		file := NewFile(pkg, filename)
		for _, t := range types {
			name := ""
			if t.Name != "" {
				name += qnaming.UpperCamelCase(t.Name)
			} else {
				name += qnaming.UpperCamelCase(string(t.BuiltInType))
			}
			file.WriteSnippet(
				Func(Var(t.BuiltInType, "v"), Var(Error, "err")).
					Named(name).
					Return(Var(t.BuiltInType)).
					Do(
						If(SnippetExpr("err != nil")).
							Do(
								Call(file.Use("log", "Panic"), Ident("err")),
							),
						Return(Ident("v")),
					),
			)
		}

		if _, err := file.Write(); err != nil {
			log.Panic(err)
		}
	}
}
