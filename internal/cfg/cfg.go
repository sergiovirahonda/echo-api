package cfg

import (
	"github.com/joeshaw/envdecode"
)

var config *Config

type (
	Config struct {
		Server
		Database
		Logger
	}
	// Server configurations
	Server struct {
		Port string `env:"SERVER_PORT,default=8080"`
		Env  string `env:"SERVER_ENV,default=local"`
	}
	// Database configurations
	Database struct {
		File string `env:"DATABASE_FILE,default=database.db"`
	}
	Logger struct {
		Level int64 `env:"LOG_LEVEL,default=4"`
	}
)

func initCfg() {
	if config != nil {
		return
	}
	config = &Config{}
	if err := envdecode.Decode(config); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	initCfg()
	return config
}
