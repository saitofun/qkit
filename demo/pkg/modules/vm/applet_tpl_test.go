package vm

import (
	"path/filepath"
	"testing"
)

var (
	conf *Config
)

func init() {
	path, err := filepath.Abs("./testdata/build/applet.yaml")
	if err != nil {
		panic(err)
	}
	conf, err = LoadConfigFrom(path)
	if err != nil {
		panic(err)
	}
}

func TestNewWasm2(t *testing.T) {

}
