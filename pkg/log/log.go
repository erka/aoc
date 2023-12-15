package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var developmentEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        zapcore.OmitKey,
	LevelKey:       "L",
	NameKey:        "N",
	CallerKey:      zapcore.OmitKey, //"C",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "M",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var (
	// defaultLogger is the default logger. It is initialized once per package
	// include upon calling DefaultLogger.
	defaultLogger     *zap.SugaredLogger
	defaultLoggerOnce sync.Once
)

// DefaultLogger returns the default logger for the package.
func DefaultLogger() *zap.SugaredLogger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLoggerFromEnv()
	})
	return defaultLogger
}

func NewLoggerFromEnv() *zap.SugaredLogger {
	level := os.Getenv("LOG_LEVEL")
	return NewLogger(level)
}

func NewLogger(level string) *zap.SugaredLogger {
	atomicLevel, _ := zap.ParseAtomicLevel(level)
	config := &zap.Config{
		Level:            atomicLevel,
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    developmentEncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	logger, err := config.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	return logger.Sugar()
}

func Debug(args ...any) {
	DefaultLogger().Debug(args...)
}
func Debugf(template string, args ...any) {
	DefaultLogger().Debugf(template, args...)
}
func Info(args ...any) {
	DefaultLogger().Debug(args...)
}
func Infof(template string, args ...any) {
	DefaultLogger().Infof(template, args...)
}
func Error(args ...any) {
	DefaultLogger().Debug(args...)
}
func Errorf(template string, args ...any) {
	DefaultLogger().Errorf(template, args...)
}
