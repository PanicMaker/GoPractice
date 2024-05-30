package Tree

import "math"

// https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		res1 := maxDepth(root.Left)
		if res1 == -1 {
			return -1
		}
		res2 := maxDepth(root.Right)
		if res2 == -1 {
			return -1
		}
		if math.Abs(float64(res1-res2)) > 1 {
			return -1
		}
		res := max(res1, res2) + 1
		return res
	}

	return maxDepth(root) != -1
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		res1 := maxDepth(root.Left)
		res2 := maxDepth(root.Right)
		return max(res1, res2) + 1
	}

	res1 := maxDepth(root.Left)
	res2 := maxDepth(root.Right)
	if math.Abs(float64(res1-res2)) > 1 {
		return false
	}

	return isBalanced2(root.Left) && isBalanced2(root.Right)
}
