package models

import (
	"os"
	"testing"

	"github.com/sergiovirahonda/echo-api/internal/cfg"
)

func TestMain(m *testing.M) {
	logger := cfg.GetLogger()
	logger.Info("Running model tests...")
	code := m.Run()
	logger.Info("Model tests finished.")
	os.Exit(code)
}
