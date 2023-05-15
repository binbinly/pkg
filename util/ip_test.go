package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceAddrs(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "InterfaceAddrs",
			args: args{v: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceAddrs(tt.args.v...)
			assert.Nil(t, err)
			t.Logf("got:%v", got)
		})
	}
}

func TestInternalIP(t *testing.T) {
	type args struct {
		dstAddr string
		network string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "InternalIP", args: args{
			dstAddr: "",
			network: "udp4",
		}},
		{name: "InternalIP-6", args: args{
			dstAddr: "[2001:4860:4860::8888]:53",
			network: "udp6",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InternalIP(tt.args.dstAddr, tt.args.network)
			t.Logf("got:%v", got)
		})
	}
}

func TestLocalIP(t *testing.T) {
	got := LocalIP()
	t.Logf("got:%v", got)
}

func TestLocalIPv4s(t *testing.T) {
	got := LocalIPv4s()
	t.Logf("got:%v", got)
}
