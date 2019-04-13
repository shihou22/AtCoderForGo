package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3}, want: 2},
		{name: "2", args: args{6}, want: 9},
		{name: "3", args: args{11}, want: 30},
		{name: "4", args: args{50}, want: 625},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.n); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{0, 0, 0, 1}, want: "-1 1 -1 0"},
		{name: "2", args: args{2, 3, 6, 6}, want: "3 10 -1 7"},
		{name: "3", args: args{31, -41, -59, 26}, want: "-126 -64 -36 -131"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveC(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3, 2}, want: 9},
		{name: "2", args: args{5, 3}, want: 1},
		{name: "3", args: args{31415, 9265}, want: 27},
		{name: "4", args: args{35897, 932}, want: 114191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveC(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solveC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveC2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3, 2}, want: 9},
		{name: "2", args: args{5, 3}, want: 1},
		{name: "3", args: args{31415, 9265}, want: 27},
		{name: "4", args: args{35897, 932}, want: 114191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveC2(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solveC2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveCX(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3, 2}, want: 9},
		{name: "2", args: args{5, 3}, want: 1},
		{name: "3", args: args{31415, 9265}, want: 27},
		{name: "4", args: args{35897, 932}, want: 114191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveCX(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solveCX() = %v, want %v", got, tt.want)
			}
		})
	}
}
