package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	InitLogger(WithLevel(InfoLevel), WithLogDir("../test/logs/"))
	Log(InfoLevel, "test info")
	Logf(InfoLevel, "test info %d", 1)

	Fields(map[string]any{"name": "app"}).Log(ErrorLevel, "test error")
	Fields(map[string]any{"name": "app"}).Logf(ErrorLevel, "test error %d", 1)

	Debug("test debug")

	Infof("test info")
	Debugf("test debug %d", 1)
}
