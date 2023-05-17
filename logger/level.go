package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
	FatalLevel = "fatal"
)

var loggerLevels = map[string]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	PanicLevel: zapcore.PanicLevel,
	FatalLevel: zapcore.FatalLevel,
}

func getZapLevel(l string) zapcore.Level {
	level, exist := loggerLevels[l]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func Debug(args ...interface{}) {
	logger.Log(DebugLevel, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Logf(DebugLevel, format, args...)
}

func Info(args ...interface{}) {
	logger.Log(InfoLevel, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf(InfoLevel, format, args...)
}

func Warn(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Logf(WarnLevel, format, args...)
}

func Error(args ...interface{}) {
	logger.Log(ErrorLevel, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Logf(ErrorLevel, format, args...)
}

func Fatal(args ...interface{}) {
	logger.Log(FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	logger.Logf(FatalLevel, format, args...)
	os.Exit(1)
}
