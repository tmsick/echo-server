package logger

import (
	"context"
	"echo-server/kontext"

	"go.uber.org/zap"
)

func New(env string, opt ...zap.Option) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)
	switch env {
	case "development":
		logger, err = zap.NewDevelopment(opt...)
	case "production":
		logger, err = zap.NewProduction(opt...)
	default:
		logger = zap.NewExample(opt...)
	}
	return logger, err
}

func WithContext(logger *zap.Logger) func(ctx context.Context) *zap.Logger {
	return func(ctx context.Context) *zap.Logger {
		return logger.WithOptions(kontext.ZapOption(ctx))
	}
}
