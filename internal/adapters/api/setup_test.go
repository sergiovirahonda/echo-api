package api

import (
	"os"
	"testing"

	"github.com/sergiovirahonda/echo-api/internal/adapters/db"
	"github.com/sergiovirahonda/echo-api/internal/cfg"
	"gorm.io/gorm"
)

var database *gorm.DB

func TestMain(m *testing.M) {
	logger := cfg.GetLogger()
	logger.Info("Running API adapters tests...")
	logger.Info("API test database...")
	database = db.NewTestConnection()
	db.Migrate(database)
	code := m.Run()
	logger.Info("API adapters tests finished.")
	os.Exit(code)
}
