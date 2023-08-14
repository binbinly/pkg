package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSortString(t *testing.T) {
	type args struct {
		m map[string]any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "TestMapSortString",
			args: args{m: map[string]any{"b": 2, "a": 1, "c": 3}},
			want: map[string]any{"a": 1, "b": 2, "c": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapSortString(tt.args.m), "MapSortString(%v)", tt.args.m)
		})
	}
}
