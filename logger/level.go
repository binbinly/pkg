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

func Debug(args ...any) {
	logger.Log(DebugLevel, args...)
}

func Debugf(format string, args ...any) {
	logger.Logf(DebugLevel, format, args...)
}

func Info(args ...any) {
	logger.Log(InfoLevel, args...)
}

func Infof(format string, args ...any) {
	logger.Logf(InfoLevel, format, args...)
}

func Warn(args ...any) {
	logger.Log(WarnLevel, args...)
}

func Warnf(format string, args ...any) {
	logger.Logf(WarnLevel, format, args...)
}

func Error(args ...any) {
	logger.Log(ErrorLevel, args...)
}

func Errorf(format string, args ...any) {
	logger.Logf(ErrorLevel, format, args...)
}

func Fatal(args ...any) {
	logger.Log(FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	logger.Logf(FatalLevel, format, args...)
	os.Exit(1)
}
