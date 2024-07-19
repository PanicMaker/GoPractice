package Tree

import (
	"slices"
)

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果先序遍历或中序遍历数组为空，则返回空节点
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 创建根节点，根据先序遍历的第一个元素确定
	root := &TreeNode{Val: preorder[0]}

	// 在中序遍历中找到根节点的位置
	// idx 可表示左子树的节点个数
	idx := slices.Index(inorder, preorder[0])

	// 递归构建左子树，左子树的先序遍历和中序遍历分别是 preorder[1:idx+1] 和 inorder[:idx]
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])

	// 递归构建右子树，右子树的先序遍历和中序遍历分别是 preorder[idx+1:] 和 inorder[idx+1:]
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	// 返回根节点
	return root
}
