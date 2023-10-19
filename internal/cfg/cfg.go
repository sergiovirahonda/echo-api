package cfg

import (
	"fmt"

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
		DatabaseUser     string `env:"DB_USER"`
		DatabasePassword string `env:"DB_PASSWORD"`
		DatabaseName     string `env:"DB_NAME"`
		DatabaseHost     string `env:"DB_HOST"`
		DatabasePort     string `env:"DB_PORT"`
		DatabaseSslMode  string `env:"DB_SSL_MODE"`
		DatabaseDsn      string
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

func (db *Config) GetDatabaseDsn() string {
	if config.Database.DatabaseSslMode == "" {
		config.Database.DatabaseSslMode = "disable"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		db.Database.DatabaseHost,
		db.Database.DatabasePort,
		db.Database.DatabaseUser,
		db.Database.DatabaseName,
		db.Database.DatabasePassword,
		db.Database.DatabaseSslMode,
	)
	return dsn

}

func (db *Config) GetTestDatabaseDsn(dbName string) string {
	// Set a default value if the environment variable is not set
	if config.Database.DatabaseSslMode == "" {
		config.Database.DatabaseSslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		db.Database.DatabaseHost,
		db.Database.DatabasePort,
		db.Database.DatabaseUser,
		dbName,
		db.Database.DatabasePassword,
		db.Database.DatabaseSslMode,
	)
	return dsn

}
