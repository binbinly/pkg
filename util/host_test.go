package util

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidIP(t *testing.T) {
	tests := []struct {
		addr   string
		expect bool
	}{
		{"127.0.0.1", false},
		{"255.255.255.255", false},
		{"0.0.0.0", false},
		{"localhost", false},
		{"10.1.0.1", true},
		{"172.16.0.1", true},
		{"192.168.1.1", true},
		{"8.8.8.8", true},
		{"1.1.1.1", true},
		{"9.255.255.255", true},
		{"10.0.0.0", true},
		{"10.255.255.255", true},
		{"11.0.0.0", true},
		{"172.15.255.255", true},
		{"172.16.0.0", true},
		{"172.16.255.255", true},
		{"172.23.18.255", true},
		{"172.31.255.255", true},
		{"172.31.0.0", true},
		{"172.32.0.0", true},
		{"192.167.255.255", true},
		{"192.168.0.0", true},
		{"192.168.255.255", true},
		{"192.169.0.0", true},
		{"fbff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"fc00::", true},
		{"fcff:1200:0:44::", true},
		{"fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"fe00::", true},
	}
	for _, test := range tests {
		t.Run(test.addr, func(t *testing.T) {
			res := isValidIP(test.addr)
			assert.Equal(t, res, test.expect)
		})
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		addr   string
		expect string
	}{
		{"127.0.0.1:80", "127.0.0.1:80"},
		{"10.0.0.1:80", "10.0.0.1:80"},
		{"172.16.0.1:80", "172.16.0.1:80"},
		{"192.168.1.1:80", "192.168.1.1:80"},
		{"0.0.0.0:80", ""},
		{"[::]:80", ""},
		{":80", ""},
	}
	for _, test := range tests {
		t.Run(test.addr, func(t *testing.T) {
			res, err := Extract(test.addr, nil)
			assert.Nil(t, err)
			if res != test.expect && (test.expect == "" && test.addr == test.expect) {
				t.Fatalf("expected %s got %s", test.expect, res)
			}
		})
	}
}

func TestExtract2(t *testing.T) {
	addr := "localhost:9001"
	lis, err := net.Listen("tcp", addr)
	assert.Nil(t, err)
	res, err := Extract(addr, lis)
	assert.Nil(t, err)
	assert.Equal(t, res, "localhost:9001")
}

func TestPort(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	assert.Nil(t, err)
	port, ok := Port(lis)
	assert.Equal(t, ok, true)
	assert.NotEqual(t, port, 0)
}

func TestExtractHostPort(t *testing.T) {
	host, port, err := ExtractHostPort("127.0.0.1:8000")
	assert.Nil(t, err)
	assert.Equal(t, host, "127.0.0.1")
	assert.Equal(t, port, uint64(8000))

	host, port, err = ExtractHostPort("www.bilibili.com:80")
	assert.Nil(t, err)
	assert.Equal(t, host, "www.bilibili.com")
	assert.Equal(t, port, uint64(80))

	host, port, err = ExtractHostPort("consul://2/33")
	assert.Error(t, err)
}
