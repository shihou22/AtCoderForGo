package main

import (
	"testing"
)

func Test_projectA(t *testing.T) {
	type args struct {
		a    int
		b    int
		c    int
		strS string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1, 2, 3, "test"}, want: "6 test"},
		{name: "2", args: args{72, 128, 256, "myonmyon"}, want: "456 myonmyon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := projectA(tt.args.a, tt.args.b, tt.args.c, tt.args.strS); got != tt.want {
				t.Errorf("projectA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc086A(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{3, 4}, want: "Even"},
		{name: "2", args: args{1, 21}, want: "Odd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc086A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("abc086A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc081A(t *testing.T) {
	type args struct {
		strS string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"101"}, want: 2},
		{name: "2", args: args{"000"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc081A(tt.args.strS); got != tt.want {
				t.Errorf("abc081A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc081B(t *testing.T) {
	type args struct {
		n  int
		wk []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{3, []int{8, 12, 40}}, want: 2},
		{name: "", args: args{4, []int{5, 6, 8, 10}}, want: 0},
		{name: "", args: args{6, []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc081B(tt.args.n, tt.args.wk); got != tt.want {
				t.Errorf("abc081B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc087B(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{2, 2, 2, 100}, want: 2},
		{name: "2", args: args{5, 1, 0, 150}, want: 0},
		{name: "3", args: args{30, 40, 50, 6000}, want: 213},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc087B(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("abc087B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc083B(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{20, 2, 5}, want: 84},
		{name: "2", args: args{10, 1, 2}, want: 13},
		{name: "3", args: args{100, 4, 16}, want: 4554},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc083B(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("abc083B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc088B(t *testing.T) {
	type args struct {
		n  int
		wk []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{2, []int{3, 1}}, want: 2},
		{name: "2", args: args{3, []int{2, 7, 4}}, want: 5},
		{name: "3", args: args{4, []int{20, 18, 2, 18}}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc088B(tt.args.n, tt.args.wk); got != tt.want {
				t.Errorf("abc088B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc085B(t *testing.T) {
	type args struct {
		n  int
		wk []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{4, []int{10, 8, 8, 6}}, want: 3},
		{name: "2", args: args{3, []int{15, 15, 15}}, want: 1},
		{name: "3", args: args{7, []int{50, 30, 50, 100, 50, 80, 30}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc085B(tt.args.n, tt.args.wk); got != tt.want {
				t.Errorf("abc085B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc085C(t *testing.T) {
	type args struct {
		n int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{9, 45000}, want: "0 9 0"},
		{name: "2", args: args{20, 196000}, want: "-1 -1 -1"},
		{name: "3", args: args{2000, 20000000}, want: "2000 0 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc085C(tt.args.n, tt.args.y); got != tt.want {
				t.Errorf("abc085C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc049C(t *testing.T) {
	type args struct {
		wk string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"erasedream"}, want: "YES"},
		{name: "2", args: args{"dreameraser"}, want: "YES"},
		{name: "3", args: args{"dreamerer"}, want: "NO"},
		{name: "2", args: args{"eraserdreameraser"}, want: "YES"},
		{name: "2", args: args{"erasedreamerase"}, want: "YES"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc049C(tt.args.wk); got != tt.want {
				t.Errorf("abc049C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abc086C(t *testing.T) {
	type args struct {
		n  int
		wk [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{2, [][]int{{3, 1, 2}, {6, 1, 1}}}, want: "Yes"},
		{name: "2", args: args{1, [][]int{{2, 100, 100}}}, want: "No"},
		{name: "3", args: args{1, [][]int{{5, 1, 1}, {100, 1, 1}}}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abc086C(tt.args.n, tt.args.wk); got != tt.want {
				t.Errorf("abc086C() = %v, want %v", got, tt.want)
			}
		})
	}
}
