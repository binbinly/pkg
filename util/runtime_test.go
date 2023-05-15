package util

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSafeGo(t *testing.T) {
	var (
		err   any
		trace []byte
	)
	rcb := func(e any, s []byte) {
		err = e
		trace = s
	}
	SafeGo(testFn2, rcb)
	time.Sleep(5 * time.Millisecond)
	assert.Equal(t, "fn1", err)
	assert.Equal(t, true, bytes.Contains(trace, []byte("panic")))
}

var (
	testFn1 = func() {
		panic("fn1")
	}
	testFn2 = func() {
		testFn1()
	}
)
