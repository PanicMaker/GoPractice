package Tree

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list

func flatten(root *TreeNode) {
	var flattenTree func(node *TreeNode) *TreeNode

	flattenTree = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		// 递归处理左子树和右子树
		leftTail := flattenTree(node.Left)
		rightTail := flattenTree(node.Right)

		// 将左子树插入到当前节点和右子树之间
		if leftTail != nil {
			leftTail.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}

		// 返回右子树的尾部
		if rightTail != nil {
			return rightTail
		} else if leftTail != nil {
			return leftTail
		} else {
			return node
		}
	}

	flattenTree(root)
}
