package format_test

import (
	"path"
	"runtime"
	"testing"

	"github.com/saitofun/qkit/gen/codegen/internal/format"
	"github.com/stretchr/testify/require"
)

func TestStdLibSet_Read(t *testing.T) {
	s := make(format.StdLibSet)

	s.WalkInit(path.Join(runtime.GOROOT(), "src"), "")

	tt := require.New(t)

	tt.Equal(false, s["json"])
	tt.Equal(true, s["encoding/json"])
}
