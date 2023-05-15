package util

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testString = "  test 中　文\u2728->?\n*\U0001F63A   "
	testBytes  = []byte(testString)
)

func TestS2B(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		s := RandString(64)
		expected := []byte(s)
		actual := S2B(s)
		assert.Equal(t, expected, actual)
		assert.Equal(t, len(expected), len(actual))
	}

	expected := testString
	actual := S2B(expected)
	assert.Equal(t, []byte(expected), actual)

	assert.Equal(t, true, S2B("") == nil)
	assert.Equal(t, testBytes, S2B(testString))
}

func TestB2S(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		b, err := RandBytes(64)
		assert.Nil(t, err)
		assert.Equal(t, string(b), B2S(b))
	}

	expected := testString
	actual := B2S([]byte(expected))
	assert.Equal(t, expected, actual)

	assert.Equal(t, true, B2S(nil) == "")
	assert.Equal(t, testString, B2S(testBytes))
}

func TestMustString(t *testing.T) {
	now := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)
	for _, v := range []struct {
		in  any
		out string
	}{
		{"Is string?", "Is string?"},
		{0, "0"},
		{0.005, "0.005"},
		{nil, ""},
		{true, "true"},
		{false, "false"},
		{[]byte(testString), testString},
		{[]int{0, 2, 1}, "[0 2 1]"},
		{map[string]any{"a": 0, "b": true, "C": []byte("c")}, "map[C:[99] a:0 b:true]"},
		{now, "2022-01-02 03:04:05"},
	} {
		assert.Equal(t, v.out, MustString(v.in))
	}
	assert.Equal(t, "2022-01-02T03:04:05Z", MustString(now, time.RFC3339))
}

func TestMustInt(t *testing.T) {
	for _, v := range []struct {
		in  any
		out int
	}{
		{"2", 2},
		{"  2 \n ", 2},
		{0b0010, 2},
		{10, 10},
		{0o77, 63},
		{0xff, 255},
		{-1, -1},
		{true, 1},
		{"0x", 0},
		{false, 0},
		{uint(11), 11},
		{uint64(11), 11},
		{int64(11), 11},
		{float32(11.0), 11},
		{1.005, 1},
		{nil, 0},
	} {
		assert.Equal(t, v.out, MustInt(v.in))
	}
}

func TestB64Encode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "6Kej56CBL+e8lueggX4g6aG25pu/JiM=", B64Encode(S2B("解码/编码~ 顶替&#")))
}

func TestB64UrlEncode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "6Kej56CBL-e8lueggX4g6aG25pu_JiM=", B64UrlEncode(S2B("解码/编码~ 顶替&#")))
}

func TestB64Decode(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []byte("解码/编码~ 顶替&#"), B64Decode("6Kej56CBL+e8lueggX4g6aG25pu/JiM="))
}

func TestB64UrlDecode(t *testing.T) {
	for _, v := range []struct {
		in  string
		out []byte
	}{
		{"6Kej56CBL-e8lueggX4g6aG25pu_JiM=", []byte("解码/编码~ 顶替&#")},
		{"123", nil},
	} {
		assert.Equal(t, v.out, B64UrlDecode(v.in))
	}
}

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = S2B(s)
	}
}

func BenchmarkS2BStdStringToBytes(b *testing.B) {
	s := strings.Repeat(testString, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}
