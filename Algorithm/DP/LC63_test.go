package DP

import (
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, want: 2},
		{name: "2", args: args{[][]int{{0, 1}, {0, 0}}}, want: 1},
		{name: "3", args: args{[][]int{{1, 0}}}, want: 0},
		{name: "4", args: args{[][]int{{1, 0}, {0, 0}}}, want: 0},
		{name: "5", args: args{[][]int{{0, 0}, {1, 1}, {0, 0}}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
