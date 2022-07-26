package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/conf/deploy"
	"github.com/saitofun/qkit/conf/env"
	"github.com/saitofun/qkit/x/misc/must"
	"github.com/saitofun/qkit/x/reflectx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Ctx struct {
	cmd       *cobra.Command
	name      string          // name app name
	feat      string          // feat git feature
	version   string          // version git version|git tag
	root      string          // root app root
	vars      []*env.Vars     // vars default env vars
	conf      []reflect.Value // conf config reflect.Value
	deployers map[string]deploy.Deployer
}

func New(setters ...OptSetter) *Ctx {
	c := &Ctx{}
	for _, setter := range setters {
		setter(c)
	}
	c.cmd = &cobra.Command{}
	if feat, ok := os.LookupEnv(envProjectFeat); ok {
		c.feat = feat
	}
	_ = os.Setenv(envProjectName, c.String())
	return c
}

func (c *Ctx) Conf(vs ...interface{}) {
	// XXX try to parse local.yml
	local, err := ioutil.ReadFile(filepath.Join(c.root, "./config/local.yml"))
	if err == nil {
		kv := make(map[string]string)
		if err = yaml.Unmarshal(local, &kv); err == nil {
			for k, v := range kv {
				_ = os.Setenv(k, v)
			}
		}
	}
	// XXX init all config elements
	for _, v := range vs {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			panic("should pass pointer for setting value")
		}
		must.NoError(c.scan(rv))
		must.NoError(c.marshal(rv))
		c.conf = append(c.conf, rv)
		rv = reflectx.Indirect(rv)
		if rv.Kind() == reflect.Struct {
			for i := 0; i < rv.NumField(); i++ {
				value := rv.Field(i)
				if !value.CanInterface() {
					continue
				}
				switch conf := value.Interface().(type) {
				case interface{ Init() }:
					conf.Init()
				case interface{ Init() error }:
					if err = conf.Init(); err != nil {
						panic(errors.Errorf("conf init: %v", err))
					}
				}
			}
		} else {
			if rv.CanInterface() {
				continue
			}
			switch conf := rv.Interface().(type) {
			case interface{ Init() }:
				conf.Init()
			case interface{ Init() error }:
				if err = conf.Init(); err != nil {
					panic(errors.Errorf("conf init: %v", err))
				}
			}
		}
	}
	if err = c.marshalDefault(); err != nil {
		panic(err)
	}
}

func (c *Ctx) AddCommand(name string, fn func(...string), commands ...func(*cobra.Command)) {
	cmd := &cobra.Command{Use: name}

	for i := range commands {
		commands[i](cmd)
	}

	cmd.Run = func(_ *cobra.Command, args []string) {
		fn(args...)
	}

	c.cmd.AddCommand(cmd)
}

func (c *Ctx) String() string {
	ret := c.name
	if c.feat != "" {
		ret += "--" + c.feat
	}
	if c.version != "" {
		ret += "@" + c.version
	}
	return ret
}

func (c *Ctx) Root() string { return c.root }

func (c *Ctx) Execute(fn func(...string), commands ...func(*cobra.Command)) {
	for i := range commands {
		commands[i](c.cmd)
	}
	c.cmd.Use = c.name
	c.cmd.Version = c.version
	c.cmd.Run = func(cmd *cobra.Command, args []string) {
		for i := range c.conf {
			c.log(c.conf[i])
		}
		fn(args...)
	}
	// TODO
	// for name, dpl := range c.deployers {
	// 	c.AddCommand(name, func(...string) {
	// 		if setter, ok := dpl.(types.DefaultSetter); ok {
	// 			setter.SetDefault()
	// 		}
	// 		filename := path.Join(c.root, name)
	// 		if err := dpl.Write(filename); err != nil {
	// 			panic(fmt.Errorf("init %s error: %v", name, err))
	// 		}
	// 	}, func(cmd *cobra.Command) {
	// 		cmd.Short = "init configuration for " + name
	// 	})
	// }
	if err := c.cmd.Execute(); err != nil {
		panic(err)
	}
}

func (c *Ctx) scan(rv reflect.Value) error {
	vars := env.NewVars(c.group(rv))

	if err := env.NewDecoder(vars).Decode(rv); err != nil {
		return err
	}
	c.vars = append(c.vars, vars)
	if _, err := env.NewEncoder(vars).Encode(rv); err != nil {
		return err
	}
	return nil
}

func (c *Ctx) marshal(rv reflect.Value) error {
	vars := env.LoadVarsFromEnviron(c.group(rv), os.Environ())
	if err := env.NewDecoder(vars).Decode(rv); err != nil {
		return err
	}
	return nil
}

func (c *Ctx) marshalDefault() error {
	m := map[string]string{}
	m["GOENV"] = "DEV"

	for _, vars := range c.vars {
		for _, v := range vars.Values {
			if !v.Optional {
				m[v.Key(vars.Prefix)] = v.Value
			}
		}
	}

	return WriteYamlFile(path.Join(c.root, "./config/default.yml"), m)
}

func (c *Ctx) log(rv reflect.Value) {
	vars := env.NewVars(c.group(rv))
	if _, err := env.NewEncoder(vars).Encode(rv); err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(vars.MaskBytes()))
}

func (c *Ctx) MarshalLocal() error {
	return nil
}

type Marshaller func(v interface{}) ([]byte, error)

const (
	envProjectName = "PRJ_NAME"
	envProjectFeat = "PRJ_FEAT"
)

// group returns config group name
func (c *Ctx) group(rv reflect.Value) string {
	group := rv.Elem().Type().Name()
	if rv.Elem().Type().Implements(types.RTypeNamed) {
		group = rv.Elem().Interface().(types.Named).Name()
	}
	return strings.ToUpper(strings.Replace(c.name+"__"+group, "-", "_", -1))
}
