package appendInt

import (
	"reflect"
	"testing"
)

func Test_appendInt(t *testing.T) {
	type args struct {
		arr []int
		v   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{arr: []int{1, 2, 3}, v: 4}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendInt(tt.args.arr, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
