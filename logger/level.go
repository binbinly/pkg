package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

var loggerLevels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func getZapLevel(l string) zapcore.Level {
	level, exist := loggerLevels[l]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func Debug(args ...interface{}) {
	logger.Log(zapcore.DebugLevel, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Logf(zapcore.DebugLevel, format, args...)
}

func Info(args ...interface{}) {
	logger.Log(zapcore.InfoLevel, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf(zapcore.InfoLevel, format, args...)
}

func Warn(args ...interface{}) {
	logger.Log(zapcore.WarnLevel, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Logf(zapcore.WarnLevel, format, args...)
}

func Error(args ...interface{}) {
	logger.Log(zapcore.ErrorLevel, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Logf(zapcore.ErrorLevel, format, args...)
}

func Fatal(args ...interface{}) {
	logger.Log(zapcore.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	logger.Logf(zapcore.FatalLevel, format, args...)
	os.Exit(1)
}
