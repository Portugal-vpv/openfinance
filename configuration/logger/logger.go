package logger

import (
	"go.uber.org/zap"
)

var z *zap.Logger

func NewLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	z = logger
}

func L() *zap.Logger {
	return z
}
