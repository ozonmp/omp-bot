package logger

import (
	"context"
	"github.com/opentracing/opentracing-go"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

func fromContext(ctx context.Context) *zap.SugaredLogger {
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		return attachedLogger
	}

	return globalLogger
}

//ErrorKV записать сообщение в error log
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

//WarnKV записать сообщение в warn log
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

//InfoKV записать сообщение в info log
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

//DebugKV записать сообщение в debug log
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

//FatalKV записать сообщение в fatal log
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

//AttachLogger добавить логер
func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

//FromContext получить логер из контекста
func FromContext(ctx context.Context) *zap.SugaredLogger {
	var result *zap.SugaredLogger
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		result = attachedLogger
	} else {
		result = globalLogger
	}

	jaegerSpan := opentracing.SpanFromContext(ctx)
	if jaegerSpan != nil {
		if spanCtx, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
			result = result.With("trace-id", spanCtx.TraceID())
		}
	}

	return result
}

//CloneWithLevel создать логер с уровнем
func CloneWithLevel(ctx context.Context, newLevel zapcore.Level) *zap.SugaredLogger {
	return FromContext(ctx).
		Desugar().
		WithOptions(WithLevel(zapcore.Level(newLevel))).
		Sugar()
}

//SetLogger установить логер
func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}

	globalLogger = notSugaredLogger.Sugar()
}

//InitLogger инициализация
func InitLogger(ctx context.Context, graylogPath string, loggingLevel zapcore.Level, serviceName string) (syncFn func()) {
	InfoKV(ctx, "Start logger")

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(graylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		FatalKV(ctx, "init logger failed ", "err", err)
	}

	InfoKV(ctx, "logger path", loggingLevel, "level", loggingLevel)

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	SetLogger(sugaredLogger.With(
		"service", serviceName,
	))

	return func() {
		notSugaredLogger.Sync() //nolint
	}
}

//ToLoggingLevel перевод строки в уровень логирования
func ToLoggingLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}
