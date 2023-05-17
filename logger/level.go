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
	logger.Log("debug", args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Logf("debug", format, args...)
}

func Info(args ...interface{}) {
	logger.Log("info", args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf("info", format, args...)
}

func Warn(args ...interface{}) {
	logger.Log("warn", args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Logf("warn", format, args...)
}

func Error(args ...interface{}) {
	logger.Log("error", args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Logf("error", format, args...)
}

func Fatal(args ...interface{}) {
	logger.Log("fatal", args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	logger.Logf("fatal", format, args...)
	os.Exit(1)
}
