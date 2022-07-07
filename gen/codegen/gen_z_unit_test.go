package codegen_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/saitofun/qkit/gen/codegen"
)

func CreateDemoFile() *File {
	filename := "examples/hello/hello.go"
	f := NewFile("main", filename)
	f.WriteSnippet(Func().Named("main").Do(
		Call(f.Use("fmt", "Println"), f.Value("Hello, 世界")),
		Call(f.Use("github.com/some/pkg", "Println"), f.Value("Hello World!")),
		Call(f.Use("github.com/another/pkg", "Println"), f.Value("Hello World!")),
		Call(f.Use("github.com/one_more/pkg", "Println"), f.Value("Hello World!")),

		Assign(AnonymousIdent).By(Call(f.Use("bytes", "NewBuffer"), f.Value(nil))),
	))

	return f
}

func Test_NewFile(t *testing.T) {
	f := CreateDemoFile()

	defer os.RemoveAll("examples")

	if _, err := f.Write(); err != nil {
		panic(err)
	}
	if raw, err := os.ReadFile(f.Name); err != nil {
		panic(err)
	} else {
		// NOTE: this test should always FAILED because the generated file
		// contains `time` and `version` information
		fmt.Println("*********the following is generated file content*********")
		fmt.Println(string(raw))
	}
}
