package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

type LogFields struct {
	RequestID string
	UserID    uint
	IP        string
	Method    string
	Path      string
	Status    int
	Duration  time.Duration
	Error     error
}

func NewLogger(level string) (*Logger, error) {
	config := zap.NewProductionConfig()

	// Set log level
	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	// Configure output
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	// Configure encoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = "stacktrace"

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return &Logger{logger}, nil
}

func (l *Logger) Info(ctx context.Context, msg string, fields LogFields) {
	zapFields := []zap.Field{
		zap.String("request_id", fields.RequestID),
		zap.Uint("user_id", fields.UserID),
		zap.String("ip", fields.IP),
		zap.String("method", fields.Method),
		zap.String("path", fields.Path),
		zap.Int("status", fields.Status),
		zap.Duration("duration", fields.Duration),
	}

	if fields.Error != nil {
		zapFields = append(zapFields, zap.Error(fields.Error))
	}

	l.Logger.Info(msg, zapFields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields LogFields) {
	zapFields := []zap.Field{
		zap.String("request_id", fields.RequestID),
		zap.Uint("user_id", fields.UserID),
		zap.String("ip", fields.IP),
		zap.String("method", fields.Method),
		zap.String("path", fields.Path),
		zap.Int("status", fields.Status),
		zap.Duration("duration", fields.Duration),
	}

	if fields.Error != nil {
		zapFields = append(zapFields, zap.Error(fields.Error))
	}

	l.Logger.Error(msg, zapFields...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields LogFields) {
	zapFields := []zap.Field{
		zap.String("request_id", fields.RequestID),
		zap.Uint("user_id", fields.UserID),
		zap.String("ip", fields.IP),
		zap.String("method", fields.Method),
		zap.String("path", fields.Path),
		zap.Int("status", fields.Status),
		zap.Duration("duration", fields.Duration),
	}

	if fields.Error != nil {
		zapFields = append(zapFields, zap.Error(fields.Error))
	}

	l.Logger.Debug(msg, zapFields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields LogFields) {
	zapFields := []zap.Field{
		zap.String("request_id", fields.RequestID),
		zap.Uint("user_id", fields.UserID),
		zap.String("ip", fields.IP),
		zap.String("method", fields.Method),
		zap.String("path", fields.Path),
		zap.Int("status", fields.Status),
		zap.Duration("duration", fields.Duration),
	}

	if fields.Error != nil {
		zapFields = append(zapFields, zap.Error(fields.Error))
	}

	l.Logger.Warn(msg, zapFields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields LogFields) {
	zapFields := []zap.Field{
		zap.String("request_id", fields.RequestID),
		zap.Uint("user_id", fields.UserID),
		zap.String("ip", fields.IP),
		zap.String("method", fields.Method),
		zap.String("path", fields.Path),
		zap.Int("status", fields.Status),
		zap.Duration("duration", fields.Duration),
	}

	if fields.Error != nil {
		zapFields = append(zapFields, zap.Error(fields.Error))
	}

	l.Logger.Fatal(msg, zapFields...)
}

func (l *Logger) Sync() error {
	return l.Logger.Sync()
}

// Helper functions for common fields
func String(key, value string) zap.Field {
	return zap.String(key, value)
}

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func Time(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
