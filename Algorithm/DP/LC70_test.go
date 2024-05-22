package DP

import "testing"

func Test_climbStairs1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2}, 2},
		{"2", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs1(tt.args.n); got != tt.want {
				t.Errorf("climbStairs1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2}, 2},
		{"2", args{3}, 3},
		{"3", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); got != tt.want {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2}, 2},
		{"2", args{3}, 3},
		{"3", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3() = %v, want %v", got, tt.want)
			}
		})
	}
}
