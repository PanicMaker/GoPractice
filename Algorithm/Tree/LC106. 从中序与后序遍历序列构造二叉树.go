package Tree

import (
	"slices"
)

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func buildTree106(inorder []int, postorder []int) *TreeNode {
	// 如果后序遍历数组或中序遍历数组为空，则返回空节点
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 创建根节点，根据后序遍历的最后一个元素确定
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// 在中序遍历中找到根节点的位置
	idx := slices.Index(inorder, root.Val)

	// 递归构建左子树，左子树的中序遍历和后序遍历分别是 inorder[:idx] 和 postorder[:idx]
	root.Left = buildTree106(inorder[:idx], postorder[:idx])

	// 递归构建右子树，右子树的中序遍历和后序遍历分别是 inorder[idx+1:] 和 postorder[idx:len(postorder)-1]
	root.Right = buildTree106(inorder[idx+1:], postorder[idx:len(postorder)-1])

	// 返回根节点
	return root
}
