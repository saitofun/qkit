package formatx_test

import (
	"path"
	"runtime"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/saitofun/qkit/gen/codegen/formatx"
)

func TestStdLibSet_Read(t *testing.T) {
	s := make(formatx.StdLibSet)
	s.WalkInit(path.Join(runtime.GOROOT(), "src"), "")

	NewWithT(t).Expect(s["json"]).To(BeFalse())
	NewWithT(t).Expect(s["encoding/json"]).To(BeTrue())
}
