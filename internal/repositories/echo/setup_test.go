package echo

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
	logger.Info("Running echo tests...")
	logger.Info("Instantiating test database...")
	database = db.NewTestConnection()
	db.Migrate(database)
	code := m.Run()
	logger.Info("Echo tests finished.")
	os.Exit(code)
}
