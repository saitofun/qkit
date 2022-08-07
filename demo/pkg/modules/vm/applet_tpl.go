package vm

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SpecVersion string       `yaml:"specVersion"`
	RepoAddr    string       `yaml:"repository"`
	Desc        string       `yaml:"description"`
	Schema      Schema       `yaml:"schema"`
	DataSources []DataSource `yaml:"dataSources"`
}

type Schema struct {
	File string `yml:"file"`
}

type DataSource struct {
	Kind    string  `yaml:"kind"`
	Name    string  `yaml:"name"`
	Network string  `yaml:"network"`
	Source  Source  `yaml:"source"`
	Mapping Mapping `yaml:"mapping"`
}

type Source struct {
	Address string `yaml:"address"`
	AbiName string `yaml:"abi"`
}

type Mapping struct {
	Kind          string         `yaml:"kind"`
	APIVersion    string         `yaml:"apiVersion"`
	Language      string         `yaml:"language"`
	Entities      []string       `yaml:"entities"`
	ABIs          []ABI          `yaml:"abis"`
	EventHandlers []EventHandler `yaml:"eventHandlers"`
	File          string         `yaml:"file"`
}

type ABI struct {
	Name string `yml:"name"`
	Path string `yml:"file"`
}

type EventHandler struct {
	Event   string `yml:"event"`
	Handler string `yml:"handler"`
}

func LoadConfigFrom(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(f)
	cfg := &Config{}
	if err = dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
