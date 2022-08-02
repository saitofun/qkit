package format_test

import (
	"path"
	"runtime"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/saitofun/qkit/gen/codegen/internal/format"
)

func TestStdLibSet_Read(t *testing.T) {
	s := make(format.StdLibSet)
	s.WalkInit(path.Join(runtime.GOROOT(), "src"), "")

	NewWithT(t).Expect(s["json"]).To(BeFalse())
	NewWithT(t).Expect(s["encoding/json"]).To(BeTrue())
}
