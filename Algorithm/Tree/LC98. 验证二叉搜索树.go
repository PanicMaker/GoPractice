package Tree

import "math"

// https://leetcode.cn/problems/validate-binary-search-tree/

// 先序遍历，通过判定节点值是否在范围内
func isValidBSTI(root *TreeNode) bool {

	var check func(node *TreeNode, left int, right int) bool
	check = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}

		// 左子树节点值应该都小于当前节点值，右子树节点值应该都大于当前节点值
		return left < node.Val && node.Val < right && check(node.Left, left, node.Val) && check(node.Right, node.Val, right)
	}

	return check(root, -math.MaxInt, math.MaxInt)
}

// 中序遍历, 遍历并检查结果是否为一个升序序列
func isValidBSTII(root *TreeNode) bool {

	pre := -math.MaxInt32

	var check func(node *TreeNode) bool
	check = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		// 递归遍历左子树
		if !check(node.Left) {
			return false
		}

		// 检查当前节点值是否大于上一个节点值
		if node.Val <= pre {
			return false
		}

		// 更新上一个节点值为当前节点值
		pre = node.Val

		return check(node.Right)
	}

	return check(root)
}
