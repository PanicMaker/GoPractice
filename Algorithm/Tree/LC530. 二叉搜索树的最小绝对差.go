package Tree

import (
	"math"
)

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/

// 将中序遍历的结果转换为一个数组，用双指针求最小差值
func getMinimumDifferenceI(root *TreeNode) int {
	nums := make([]int, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		nums = append(nums, node.Val)

		dfs(node.Right)
	}

	dfs(root)

	res := math.MaxInt
	for i, j := 0, 1; j < len(nums); i++ {
		res = min(res, nums[j]-nums[i])
		j++
	}

	return res
}

// 利用中序遍历的结果是一个有序数列，两两求差并与最小差值比较
func getMinimumDifferenceII(root *TreeNode) int {
	res := math.MaxInt
	var pre *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if pre != nil {
			res = min(res, node.Val-pre.Val)
		}

		pre = node

		dfs(node.Right)
	}

	dfs(root)

	return res
}
