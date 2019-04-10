package main

import "testing"

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
