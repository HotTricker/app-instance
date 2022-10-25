package miniapp

import (
	"github.com/Unknwon/goconfig"
)

type (
	Config struct {
		Log    *LogConfig
		Server *ServerConfig
	}

	LogConfig struct {
		Path string
	}

	ServerConfig struct {
		Addr string
	}
)

var (
	configFile *goconfig.ConfigFile
)

func loadConfig(path string) error {
	var err error
	configFile, err = goconfig.LoadConfigFile(path)
	if err != nil {
		return err
	}

	App.cfg = &Config{
		Log: &LogConfig{
			Path: configOrDefault("log", "path", "stdout"),
		},
		Server: &ServerConfig{
			Addr: configOrDefault("serve", "addr", "8878"),
		},
	}

	return nil
}

func configOrDefault(section, key, useDefault string) string {
	val, err := configFile.GetValue(section, key)
	if err != nil {
		return useDefault
	}
	return val
}

func configIntOrDefault(section, key string, useDefault int) int {
	val, err := configFile.Int(section, key)
	if err != nil {
		return useDefault
	}
	return val
}
