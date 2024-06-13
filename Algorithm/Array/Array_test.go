package Array

import (
	"reflect"
	"testing"
)

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
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if got := tt.args.nums1; !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("After merge %v, expect %v", got, tt.expect)
			}
		})
	}
}
