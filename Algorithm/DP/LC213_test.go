package DP

import "testing"

func Test_robII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 1}}, 4},
		{"2", args{[]int{2, 3, 2}}, 3},
		{"3", args{[]int{1, 2, 3}}, 3},
		{"4", args{[]int{0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robII(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
