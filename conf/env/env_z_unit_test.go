package env_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/saitofun/qlib/container/qptr"
	"github.com/sincospro/qkit/base/types"
	. "github.com/sincospro/qkit/conf/env"
)

func TestPathWalker(t *testing.T) {
	p := NewPathWalker()

	p.Enter("parent")
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent"}))
	NewWithT(t).Expect(p.String()).To(Equal("parent"))

	p.Enter("child")
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent", "child"}))
	NewWithT(t).Expect(p.String()).To(Equal("parent_child"))

	p.Enter("prop")
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent", "child", "prop"}))
	NewWithT(t).Expect(p.String()).To(Equal("parent_child_prop"))

	p.Enter(100)
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent", "child", "prop", 100}))
	NewWithT(t).Expect(p.String()).To(Equal("parent_child_prop_100"))

	p.Exit()
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent", "child", "prop"}))
	NewWithT(t).Expect(p.String()).To(Equal("parent_child_prop"))

	p.Exit()
	NewWithT(t).Expect(p.Paths()).To(Equal([]interface{}{"parent", "child"}))
	NewWithT(t).Expect(p.String()).To(Equal("parent_child"))
}

type SubConfig struct {
	Duration     types.Duration
	Password     types.Password `env:""`
	Key          string         `env:""`
	Bool         bool
	Map          map[string]string
	Func         func() error
	defaultValue bool
}

func (c *SubConfig) SetDefault() {
	c.defaultValue = true
}

type Config struct {
	Map       map[string]string
	Slice     []string `env:""`
	PtrString *string  `env:""`
	Host      string   `env:""`
	SubConfig
	Config SubConfig
}

func TestEnvVars(t *testing.T) {
	c := Config{}

	c.Duration = types.Duration(10 * time.Second)
	c.Password = types.Password("123123")
	c.Key = "123456"
	c.PtrString = qptr.String("123456=")
	c.Slice = []string{"1", "2"}
	c.Config.Key = "key"
	c.Config.defaultValue = true
	c.defaultValue = true

	envVars := NewVars("S")

	t.Run("Encoding", func(t *testing.T) {
		data, _ := NewEncoder(envVars).Encode(&c)
		NewWithT(t).Expect(string(data)).To(
			Equal(`S__Bool=false
S__Config_Bool=false
S__Config_Duration=0s
S__Config_Key=key
S__Config_Password=
S__Duration=10s
S__Host=
S__Key=123456
S__Password=123123
S__PtrString=123456=
S__Slice_0=1
S__Slice_1=2
`))
	})

	t.Run("Decoding", func(t *testing.T) {
		data, _ := NewEncoder(envVars).Encode(&c)

		envVars := LoadVarsFromEnviron("S", strings.Split(string(data), "\n"))

		c2 := Config{}
		err := NewDecoder(envVars).Decode(&c2)

		NewWithT(t).Expect(err).To(BeNil())
		NewWithT(t).Expect(c2).To(Equal(c))
	})
}
