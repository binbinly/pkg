package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	type args struct {
		ss []int
		s  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "InSlice-true", args: args{
			ss: []int{1, 2, 3},
			s:  1,
		}, want: true},
		{name: "InSlice-false", args: args{
			ss: []int{1, 2, 3},
			s:  4,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, InSliceInt(tt.args.s, tt.args.ss), "InSlice(%v, %v)", tt.args.ss, tt.args.s)
		})
	}
}

func TestSliceBigFilter(t *testing.T) {
	type args struct {
		a []int
		f func(v int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "SliceBigFilter", args: args{
			a: []int{1, 2, 3},
			f: func(v int) bool {
				return v > 2
			},
		}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceBigFilter(tt.args.a, tt.args.f), "SliceBigFilter(%v)", tt.args.a)
		})
	}
}

func TestSliceDeleteElem(t *testing.T) {
	type args struct {
		i int
		s []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "SliceDeleteElem", args: args{
			i: 1,
			s: []any{1, 2, 3},
		}, want: []any{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := SliceDeleteElem(tt.args.i, tt.args.s)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSliceIntJoin(t *testing.T) {
	type args struct {
		s   []int
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "SliceIntJoin", args: args{
			s:   []int{1, 2, 3},
			sep: "-",
		}, want: "1-2-3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceIntJoin(tt.args.s, tt.args.sep), "SliceIntJoin(%v, %v)", tt.args.s, tt.args.sep)
		})
	}
}

func TestSliceReverse(t *testing.T) {
	type args struct {
		s []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "SliceReverse", args: args{s: []any{1, 2, 3}}, want: []any{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceReverse(tt.args.s)
			assert.Equalf(t, tt.want, tt.args.s, "SliceReverse(%v)", tt.args.s)
		})
	}
}

func TestSliceShuffle(t *testing.T) {
	type args struct {
		s []any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "SliceShuffle", args: args{s: []any{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceShuffle(tt.args.s)
			t.Logf("shuffle:%v", tt.args.s)
		})
	}
}

func TestSliceSmallFilter(t *testing.T) {
	type args struct {
		a []int
		f func(v int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "SliceSmallFilter", args: args{
			a: []int{1, 2, 3},
			f: func(v int) bool {
				return v > 2
			},
		}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceSmallFilter(tt.args.a, tt.args.f), "SliceSmallFilter(%v)", tt.args.a)
		})
	}
}

func TestSliceToInt(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name   string
		args   args
		wantIi []int
	}{
		{name: "", args: args{ss: []string{"1", "2", "a", "123a"}}, wantIi: []int{1, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantIi, SliceToInt(tt.args.ss), "SliceToInt(%v)", tt.args.ss)
		})
	}
}

func TestSliceIntDeduplication(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestSliceIntDeduplication",
			args: args{a: []int{1, 1, 2, 2, 3, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceIntDeduplication(tt.args.a), "SliceIntDeduplication(%v)", tt.args.a)
		})
	}
}
