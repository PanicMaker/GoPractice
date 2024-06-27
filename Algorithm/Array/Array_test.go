package Array

import (
	"reflect"
	"testing"
)

func TestLC15(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name   string
		args   args
		expect [][]int
	}{
		{"1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{"2", args{[]int{0, 1, 1}}, [][]int{}},
		{"3", args{[]int{0, 0, 0}}, [][]int{[]int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumII(tt.args.nums); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("threeSum() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestLC88(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	var tests = []struct {
		name   string
		args   args
		expect []int
	}{
		{"1", args{[]int{1}, 1, []int{}, 0}, []int{1}},
		{"2", args{make([]int, 1), 0, []int{1}, 1}, []int{1}},
		{"3", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if got := tt.args.nums1; !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("After merge %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestLC80(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name   string
		args   args
		expect int
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"2", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
		{"3", args{[]int{1}}, 1},
		{"4", args{[]int{1, 1}}, 2},
		{"5", args{[]int{1, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates80II(tt.args.nums); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("removeDuplicates80() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestLC169(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name   string
		args   args
		expect int
	}{
		{"1", args{[]int{3, 2, 3}}, 3},
		{"2", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementIII(tt.args.nums); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("majorityElement() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestLC189(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	var tests = []struct {
		name   string
		args   args
		expect []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"2", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateIII(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.expect) {
				t.Errorf("rotateI() = %v, expect %v", tt.args.nums, tt.expect)
			}
		})
	}
}

func TestLC167(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"1", args{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{"1", args{[]int{-1, 0}, -1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC238(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name   string
		args   args
		expect []int
	}{
		{"1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"2", args{[]int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelfII(tt.args.nums); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("productExceptSelfI() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestLC134(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	var tests = []struct {
		name   string
		args   args
		expect int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}}, 3},
		{"2", args{[]int{2, 3, 4}, []int{3, 4, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("v() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
