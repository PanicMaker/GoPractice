package LinkedList

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates82(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{head: CreateLinkedList([]int{1, 2, 3, 3, 4, 4, 5})}, CreateLinkedList([]int{1, 2, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates82(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates82() = %v, want %v", got, tt.want)
			}
		})
	}
}
