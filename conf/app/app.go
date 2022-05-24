package app

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/sincospro/qkit/conf/deploy"
	"gopkg.in/yaml.v2"
)

type OptSetter = func(conf *Ctx)

func WithName(name string) OptSetter { return func(c *Ctx) { c.name = name } }

func WithVersion(version string) OptSetter { return func(c *Ctx) { c.version = version } }

func WithFeat(feat string) OptSetter { return func(c *Ctx) { c.feat = feat } }

func WithRoot(root string) OptSetter {
	_, filename, _, _ := runtime.Caller(1)
	return func(c *Ctx) { c.root = filepath.Join(filepath.Dir(filename), root) }
}

func WithDeployer(deployers ...deploy.Deployer) OptSetter {
	return func(c *Ctx) {
		if c.deployers == nil {
			c.deployers = make(map[string]deploy.Deployer)
		}
		for _, dpl := range deployers {
			c.deployers[dpl.Name()] = dpl
		}
	}
}

func WriteYamlFile(filename string, v interface{}) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	root := filepath.Dir(filename)
	if root != "" {
		if err = os.MkdirAll(root, os.ModePerm); err != nil {
			return err
		}
	}
	return os.WriteFile(filename, data, os.ModePerm)
}
