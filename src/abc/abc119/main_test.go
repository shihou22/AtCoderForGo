package main

import "testing"

func Test_solvoA(t *testing.T) {
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
			if got := solvoA(tt.args.dateS); got != tt.want {
				t.Errorf("solvoA() = %v, want %v", got, tt.want)
			}
		})
	}
}
