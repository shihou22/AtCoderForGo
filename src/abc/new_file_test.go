package main

import "testing"

func Test_hoge(t *testing.T) {
	type args struct {
		x int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{x: 1000, a: 108, b: 108}, want: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hoge(tt.args.x, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
