package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{5, 2}, want: 1},
		{name: "2", args: args{6, 4}, want: 2},
		{name: "3", args: args{90, 30}, want: 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{26}, want: 1},
		{name: "2", args: args{41}, want: 4},
		{name: "3", args: args{100000}, want: 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.n); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveC(t *testing.T) {
	type args struct {
		n  int
		aA []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{4, []int{100, 150, 130, 120}}, want: 40},
		{name: "2", args: args{4, []int{100, 125, 80, 110}}, want: 40},
		{name: "3", args: args{9, []int{314, 159, 265, 358, 979, 323, 846, 264, 338}}, want: 310},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveC(tt.args.n, tt.args.aA); got != tt.want {
				t.Errorf("solveC() = %v, want %v", got, tt.want)
			}
		})
	}
}
