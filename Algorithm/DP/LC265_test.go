package DP

import "testing"

func Test_minCostII(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{costs: [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}}, 10},
		{"2", args{costs: [][]int{{1, 5, 4}, {2, 9, 4}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostII(tt.args.costs); got != tt.want {
				t.Errorf("minCostII() = %v, want %v", got, tt.want)
			}
		})
	}
}
