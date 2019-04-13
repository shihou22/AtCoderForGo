package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		numH int
		numW int
		h    int
		w    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{numH: 3, numW: 2, h: 2, w: 1}, want: 1},
		{name: "2", args: args{numH: 5, numW: 5, h: 2, w: 3}, want: 6},
		{name: "3", args: args{numH: 2, numW: 4, h: 2, w: 4}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.numH, tt.args.numW, tt.args.h, tt.args.w); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		bArray []int
		aArray [][]int
		n      int
		m      int
		c      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 2, 3}, [][]int{{3, 2, 1}, {1, 2, 2}}, 2, 3, -10}, want: 1},
		{name: "2", args: args{[]int{-2, 5}, [][]int{{100, 41}, {100, 40}, {-3, 0}, {-6, -2}, {18, -13}}, 5, 2, -4}, want: 2},
		{name: "3", args: args{[]int{100, -100, 0}, [][]int{{0, 100, 100}, {100, 100, 100}, {-100, 100, 100}}, 2, 3, -10}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.bArray, tt.args.aArray, tt.args.n, tt.args.m, tt.args.c); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mainBInputTest(t *testing.T) {
	type args struct {
		bArrayArg []int
		aArrayArg [][]int
		n         int
		m         int
		c         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 2, 3}, [][]int{{3, 2, 1}, {1, 2, 2}}, 2, 3, -10}, want: 1},
		{name: "2", args: args{[]int{-2, 5}, [][]int{{100, 41}, {100, 40}, {-3, 0}, {-6, -2}, {18, -13}}, 5, 2, -4}, want: 2},
		{name: "3", args: args{[]int{100, -100, 0}, [][]int{{0, 100, 100}, {100, 100, 100}, {-100, 100, 100}}, 2, 3, -10}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mainBInputTest(tt.args.bArrayArg, tt.args.aArrayArg, tt.args.n, tt.args.m, tt.args.c); got != tt.want {
				t.Errorf("mainBInputTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
