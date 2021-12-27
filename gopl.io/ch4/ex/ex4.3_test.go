package ex

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s *[4]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{&[4]int{1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.s)
			fmt.Println(tt.args.s)
		})
	}
}
