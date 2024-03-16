package DP

import (
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 7}, 28},
		{"2", args{3, 2}, 3},
		{"3", args{7, 3}, 28},
		{"4", args{3, 3}, 6},
		{"5", args{1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
