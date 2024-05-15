package Tree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{root: &TreeNode{Val: 1, Left: nil, Right: nil}}, 1},
		{"2", args{root: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth1(tt.args.root); got != tt.want {
				t.Errorf("maxDepth1() = %v, want %v", got, tt.want)
			}
			if got := maxDepth2(tt.args.root); got != tt.want {
				t.Errorf("maxDepth1() = %v, want %v", got, tt.want)
			}
		})
	}
}
