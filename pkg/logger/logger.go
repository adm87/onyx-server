package logger

import (
	"errors"

	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *config.LoggerConfig) (*zap.Logger, error) {
	lvl, err := parseLogLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	zapCfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(lvl),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func parseLogLevel(level string) (zapcore.Level, error) {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		return 0, errors.New("invalid log level: " + level)
	}
	return zapLevel, nil
}
