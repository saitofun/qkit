package validator_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	. "github.com/saitofun/qkit/kit/validator"
	. "github.com/saitofun/qkit/kit/validator/strfmt"
)

func ExampleNewRegexpStrfmtValidator() {
	fmt.Println(AlphaValidator.Validate("a"))
	fmt.Println(AlphaValidator.Validate("1"))
	fmt.Println(EmailValidator.Validate("a.b.c+123@xxx.com"))
	// Output:
	// <nil>
	// alpha ^[a-zA-Z]+$ not match 1
	// <nil>
}

func TestStrfmtValidator_Validate(t *testing.T) {
	cases := []struct {
		value interface{}
		rule  string
		vldt  *StrFmt
	}{
		{"abc", "@alpha", AlphaValidator},
		{"a.b.c+123@xxx.com", "@email", EmailValidator},
	}

	for _, c := range cases {
		name := fmt.Sprintf(
			"%s|%s|%s", c.vldt.Names()[0], c.rule, c.value,
		)
		t.Run(name, func(t *testing.T) {
			fact := NewFactory()
			fact.Register(c.vldt)
			v, err := fact.Compile(bg, []byte(c.rule), rttString)
			NewWithT(t).Expect(err).To(BeNil())
			err = v.Validate(c.value)
			NewWithT(t).Expect(err).To(BeNil())
		})
	}
}

func TestStrfmtValidator_ValidateFailed(t *testing.T) {
	cases := []struct {
		value interface{}
		rule  string
		vldt  *StrFmt
	}{
		{1, "@number", NumberValidator},
		{".", "@number", NumberValidator},
		{"x#abc.com", "@email", EmailValidator},
		{"123", "@alpha", AlphaValidator},
	}

	for _, c := range cases {
		name := fmt.Sprintf(
			"%s|%s|%s", c.vldt.Names()[0], c.rule, c.value,
		)
		t.Run(name, func(t *testing.T) {
			err := c.vldt.Validate(c.value)
			NewWithT(t).Expect(err).NotTo(BeNil())
			// t.Logf("\n%v", err)
		})
	}
}
