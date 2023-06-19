package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatAmount(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{amount: 1.16}, want: 116},
		{name: "2", args: args{amount: 1.13}, want: 113},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatAmount(tt.args.amount), "FormatAmount(%v)", tt.args.amount)
		})
	}
}

func TestFormatResUrl(t *testing.T) {
	type args struct {
		dfs string
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"dfs", "url"}, want: "dfs/url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatResUrl(tt.args.dfs, tt.args.url), "FormatResUrl(%v, %v)", tt.args.dfs, tt.args.url)
		})
	}
}

func TestParseAmount(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{amount: 116}, want: 1.16},
		{name: "2", args: args{amount: 101}, want: 1.01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseAmount(tt.args.amount), "ParseAmount(%v)", tt.args.amount)
		})
	}
}
