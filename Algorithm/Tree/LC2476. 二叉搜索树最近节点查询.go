package Tree

import (
	"sort"
)

// https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/

func closestNodes(root *TreeNode, queries []int) [][]int {
	res := make([][]int, 0)
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

	for _, v := range queries {
		tmp := make([]int, 2)
		tmp[0] = -1
		tmp[1] = -1

		// 用双循环找两边会超时
		// // Find closest smaller node
		// for i := len(nums) - 1; i >= 0; i-- {
		// 	if nums[i] <= v {
		// 		tmp[0] = nums[i]
		// 		break
		// 	}
		// }
		//
		// // Find closest larger node
		// for i := 0; i < len(nums); i++ {
		// 	if nums[i] >= v {
		// 		tmp[1] = nums[i]
		// 		break
		// 	}
		// }

		// Find closest smaller node using binary search
		idxSmaller := sort.Search(len(nums), func(i int) bool { return nums[i] >= v })
		if idxSmaller > 0 {
			tmp[0] = nums[idxSmaller-1]
		}

		// Find closest larger node using binary search
		idxLarger := sort.Search(len(nums), func(i int) bool { return nums[i] > v })
		if idxLarger < len(nums) {
			tmp[1] = nums[idxLarger]
		}

		res = append(res, tmp)
	}

	return res
}
