package Tree

import (
	"slices"
)

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	leftVal := preorder[1]
	leftSize := slices.Index(postorder, leftVal) + 1
	left := constructFromPrePost(preorder[1:leftSize+1], postorder[:leftSize])
	right := constructFromPrePost(preorder[leftSize+1:], postorder[leftSize:len(postorder)-1])

	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}
