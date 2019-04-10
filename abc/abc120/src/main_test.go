package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
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
		{name: "1", args: args{2, 11, 4}, want: 4},
		{name: "2", args: args{3, 9, 5}, want: 3},
		{name: "3", args: args{100, 1, 10}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		a int
		b int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{a: 8, b: 12, k: 2}, want: 2},
		{name: "2", args: args{a: 100, b: 50, k: 4}, want: 5},
		{name: "3", args: args{a: 1, b: 1, k: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.a, tt.args.b, tt.args.k); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}
