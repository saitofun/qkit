package statusxgen_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/saitofun/qkit/kit/statusxgen"
	"github.com/saitofun/qkit/x/pkgx"
)

func TestGenerator(t *testing.T) {
	cwd, _ := os.Getwd()
	pkg, _ := pkgx.LoadFrom(filepath.Join(cwd, "./__examples__"))

	g := statusxgen.New(pkg)

	g.Scan("StatusError")
	g.Output(cwd)
}
