package deps

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/saitofun/qlib/container/qtype"
)

type ConfigLoader struct {
	Path   string
	Sep    byte
	Name   string
	Values map[string]string
	loaded *qtype.Bool
}

func (c *ConfigLoader) Load() error {
	f, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewReader(f)

	if c.Values == nil {
		c.Values = make(map[string]string)
	}

	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = bytes.TrimSpace(line)
		// seperated line
		if len(line) == 0 {
			continue
		}
		// comments
		if line[0] == '#' {
			continue
		}
		// config name
		if line[0] == '[' && line[len(line)-1] == ']' {
			c.Name = string(line[1 : len(line)-2])
			continue
		}
		kv := bytes.SplitN(line, []byte{c.Sep}, 2)
		if len(kv) < 2 {
			continue
		}
		kv[0] = bytes.TrimSpace(kv[0])
		kv[1] = bytes.TrimSpace(kv[1])
		c.Values[string(kv[0])] = string(kv[1])
	}
	return nil
}

func (c *ConfigLoader) Into(v interface{}) error {
	return nil
}

func (c *ConfigLoader) OverWrite(k, v string) {}
