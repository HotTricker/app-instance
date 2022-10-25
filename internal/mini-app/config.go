package miniapp

import (
	"app-instance/internal/pkg/godb"

	"github.com/Unknwon/goconfig"
)

type (
	Config struct {
		Log    *LogConfig
		Server *ServerConfig
		Db     *godb.DbConfig
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
		Db: &godb.DbConfig{
			Unix:            configOrDefault("database", "unix", ""),
			Host:            configOrDefault("database", "host", "localhost"),
			Port:            configIntOrDefault("database", "port", 3306),
			Charset:         "utf8",
			User:            configOrDefault("database", "user", "root"),
			Pass:            configOrDefault("database", "password", "123456"),
			DbName:          configOrDefault("database", "dbname", "test"),
			TablePrefix:     "",
			MaxIdleConns:    configIntOrDefault("database", "max_idle_conns", 100),
			MaxOpenConns:    configIntOrDefault("database", "max_open_conns", 200),
			ConnMaxLifeTime: configIntOrDefault("database", "conn_max_life_time", 500),
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
