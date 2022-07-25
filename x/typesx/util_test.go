package typesx_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/saitofun/qkit/x/typesx"
)

func TestTypeFor(t *testing.T) {
	cases := []string{
		"string",
		"int",
		"map[int]int",
		"[]int",
		"[2]int",
		"error",
	}
	for _, c := range cases {
		NewWithT(t).Expect(typesx.FromGoType(typesx.TypeFor(c).String())).To(Equal(c))
	}
}
