package main

import "testing"

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
