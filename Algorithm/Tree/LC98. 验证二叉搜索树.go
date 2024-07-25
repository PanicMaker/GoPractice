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

// 后序遍历
func isValidBSTIII(root *TreeNode) bool {

	var check func(node *TreeNode) (int, int)
	check = func(node *TreeNode) (int, int) {
		// 如果当前节点为空，返回最大整数和最小整数。这表示空子树对其父节点的判断没有任何影响
		if node == nil {
			return math.MaxInt, math.MinInt
		}

		lMin, lMax := check(node.Left)
		rMin, rMax := check(node.Right)

		// 如果当前节点的值小于等于左子树的最大值，或者大于等于右子树的最小值，那么当前子树不是BST
		// 返回 math.MinInt 和 math.MaxInt 表示这个子树无效
		if node.Val <= lMax || node.Val >= rMin {
			return math.MinInt, math.MaxInt
		}

		// 返回当前子树的最小值和最大值，最小值是左子树的最小值和当前节点值中的较小者，最大值是右子树的最大值和当前节点值中的较大者
		return min(lMin, node.Val), max(rMax, node.Val)
	}

	_, mx := check(root)

	// 如果最大值不是 math.MaxInt，则树是一个有效的二叉搜索树
	return mx != math.MaxInt
}
