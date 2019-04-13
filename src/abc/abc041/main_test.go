package main

import (
	"reflect"
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		s    string
		numI int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"atcoder", 3}, want: "c"},
		{name: "2", args: args{"beginner", 1}, want: "b"},
		{name: "3", args: args{"contest", 7}, want: "t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.s, tt.args.numI); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		a int64
		b int64
		c int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "", args: args{2, 3, 4}, want: 24},
		{name: "", args: args{10000, 1000, 100}, want: 1000000000},
		{name: "", args: args{100000, 1, 100000}, want: 999999937},
		{name: "", args: args{1000000000, 1000000000, 1000000000}, want: 999999664},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveC(t *testing.T) {
	type args struct {
		n  int
		wk []StudentC
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{n: 3, wk: []StudentC{{Index: 1, Height: 140}, {Index: 2, Height: 180}, {Index: 3, Height: 160}}}, want: []string{"2", "3", "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveC(tt.args.n, tt.args.wk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveC() = %v, want %v", got, tt.want)
			}
		})
	}
}
