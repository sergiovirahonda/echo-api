package dto

import (
	"os"
	"testing"

	"github.com/sergiovirahonda/echo-api/internal/cfg"
)

func TestMain(m *testing.M) {
	logger := cfg.GetLogger()
	logger.Info("Running infrastructure tests...")
	code := m.Run()
	logger.Info("Infrastructure tests finished.")
	os.Exit(code)
}
