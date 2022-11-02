package cmd_test

import (
	"testing"

	"golang.org/x/mod/modfile"

	"github.com/saitofun/qkit/gen/cmd/toolkit/cmd"
)

func TestModulePath(t *testing.T) {
	mod := modfile.ModulePath([]byte(`
module github.com/saitofun/qkit

go 1.18
`))
	t.Logf(mod)

	mod = modfile.ModulePath([]byte(`xxx`))
	t.Logf(mod)
}

func TestPrepareArg(t *testing.T) {
	cmd.PrepareArgs()
}

func TestFormatRoot(t *testing.T) {
	_ = cmd.FormatRoot("golang.org/x", "./testdata")
}
