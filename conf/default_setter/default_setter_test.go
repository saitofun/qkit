package default_setter_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/saitofun/qkit/conf/default_setter"
	"github.com/saitofun/qkit/x/ptrx"
)

func TestStruct(t *testing.T) {
	type A struct {
		A int
		B float32
		C *string
		d string
	}
	dft := A{1, 2, ptrx.String("abc"), "def"}
	tar := A{}
	NewWithT(t).Expect(default_setter.Set(dft, &tar)).To(BeNil())
	NewWithT(t).Expect(dft).To(Equal(tar))
}
