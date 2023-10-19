package cfg

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	conf   = GetConfig()
	logger *zap.SugaredLogger
)

func initLog() {
	if logger != nil {
		return
	}
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, getLevel(conf.Logger.Level))
	defaultLogger := zap.New(core, zap.AddCaller())
	logger = defaultLogger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	initLog()
	return logger
}

func getLevel(lvl int64) zapcore.Level {
	switch lvl {
	case 0:
		return zap.InfoLevel
	case 1:
		return zap.WarnLevel
	case 2:
		return zap.ErrorLevel
	}
	return zap.DebugLevel
}
