package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Log("info", "test info")
	Logf("info", "test info %d", 1)

	Fields(map[string]any{"name": "app"}).Log("error", "test error")
	Fields(map[string]any{"name": "app"}).Logf("error", "test error %d", 1)

	Debug("test error")
	Debugf("test error %d", 1)
}
