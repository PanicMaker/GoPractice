package DP

import (
	"testing"
)

func TestLC45(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 1, 1, 4}}, 2},
		{"2", args{[]int{2, 3, 0, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpI(tt.args.nums); got != tt.want {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC55(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"2", args{[]int{3, 2, 1, 0, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJumpII(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC673(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 4, 7}}, 2},
		{"2", args{[]int{2, 2, 2, 2, 2}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
