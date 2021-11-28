package logger

import (
	"context"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/rs/zerolog"
	"os"
)

type loggerKeyType uint8

const loggerKey loggerKeyType = 1

type logger struct {
	log *zerolog.Logger
}

func NewLoggerWithContext(ctx context.Context, cfg *config.Config) (*logger, context.Context) {
	level := zerolog.InfoLevel
	if cfg != nil {
		level = levelFromString(cfg.Logger.LogLevel)
	}

	log := zerolog.New(os.Stdout).Level(level)
	resLog := &logger{log: &log}

	resCtx := context.WithValue(ctx, loggerKey, resLog)
	return resLog, resCtx
}

func NewLogger(cfg *config.Config) *logger {
	level := zerolog.InfoLevel
	if cfg != nil {
		level = levelFromString(cfg.Logger.LogLevel)
	}

	log := zerolog.New(os.Stdout).Level(level)

	return &logger{log: &log}
}

func LoggerFromContext(ctx context.Context) *logger {
	logFromCtx := ctx.Value(loggerKey)
	log, ok := logFromCtx.(*logger)
	if !ok {
		log = NewLogger(nil)
	}

	return log
}

func (l logger) Info() *zerolog.Event {
	return l.log.Info()
}

func (l logger) Error() *zerolog.Event {
	return l.log.Error()
}

func (l logger) Debug() *zerolog.Event {
	return l.log.Debug()
}

func (l logger) Warn() *zerolog.Event {
	return l.log.Warn()
}

func (l logger) Fatal() *zerolog.Event {
	return l.log.Fatal()
}
