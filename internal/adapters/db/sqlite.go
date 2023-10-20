package db

import (
	"github.com/glebarez/sqlite"
	"github.com/sergiovirahonda/echo-api/internal/cfg"
	dto "github.com/sergiovirahonda/echo-api/internal/infrastructure/dtos"
	"gorm.io/gorm"
)

func NewConnection(config cfg.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewTestConnection() *gorm.DB {
	// In memory database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	logger := cfg.GetLogger()
	err := db.AutoMigrate(&dto.Echo{})
	logger.Info("Migrations executed")
	if err != nil {
		panic(err)
	}
}
