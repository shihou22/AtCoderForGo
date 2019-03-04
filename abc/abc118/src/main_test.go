package main

import (
	"testing"
)

func Test_solveA(t *testing.T) {
	type args struct {
		dateS string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"2019/04/30"}, want: "Heisei"},
		{name: "2", args: args{"2019/11/01"}, want: "TBD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveA(tt.args.dateS); got != tt.want {
				t.Errorf("solveA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveB(t *testing.T) {
	type args struct {
		a []float64
		b []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{a: []float64{10000, 0.10000000}, b: []string{"JPY", "BTC"}}, want: 48000.0},
		{name: "2", args: args{a: []float64{100000000, 100.00000000, 0.00000001}, b: []string{"JPY", "BTC", "BTC"}}, want: 138000000.0038},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveB(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveB() = %v, want %v", got, tt.want)
			}
		})
	}
}
