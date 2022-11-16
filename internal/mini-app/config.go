package miniapp

import (
	"app-instance/internal/pkg/godb"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type (
	Config struct {
		Log    *LogConfig     `yaml:"log"`
		Server *ServerConfig  `yaml:"server"`
		Db     *godb.DbConfig `yaml:"database"`
	}

	LogConfig struct {
		Path string `yaml:"path"`
	}

	ServerConfig struct {
		Addr string `yaml:"addr"`
	}
)

func loadConfig(path string) error {
	txt, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(txt, cfg)
	if err != nil {
		return err
	}

	App.cfg = cfg

	return nil
}
