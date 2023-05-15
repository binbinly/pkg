package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	Log(zapcore.InfoLevel, "test info")
	Logf(zapcore.InfoLevel, "test info %d", 1)

	Fields(map[string]any{"name": "app"}).Log(zapcore.ErrorLevel, "test error")
	Fields(map[string]any{"name": "app"}).Logf(zapcore.ErrorLevel, "test error %d", 1)

	Debug("test error")
	Debugf("test error %d", 1)
}
