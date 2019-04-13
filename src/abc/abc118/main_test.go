package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{a: 4, b: 12}, want: 16},
		{name: "2", args: args{a: 8, b: 20}, want: 12},
		{name: "3", args: args{a: 1, b: 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		n  int
		m  int
		kA []int
		aA [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3, 4, []int{2, 3, 2}, [][]int{{1, 3}, {1, 2, 3}, {3, 2}}}, want: 1},
		{name: "2", args: args{5, 5, []int{4, 4, 4, 4, 4}, [][]int{{2, 3, 4, 5}, {1, 3, 4, 5}, {1, 2, 4, 5}, {1, 2, 3, 5}, {1, 2, 3, 4}}}, want: 0},
		{name: "3", args: args{1, 30, []int{3}, [][]int{{5, 10, 30}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.n, tt.args.m, tt.args.kA, tt.args.aA); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}
