package Tree

import (
	"math"
)

// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/

func goodNodes(root *TreeNode) int {
	res := 0
	// nodes := make([]int, 0)

	var dfs func(node *TreeNode, maxVal int)
	dfs = func(node *TreeNode, maxVal int) {
		if node == nil {
			return
		}

		maxVal = max(maxVal, node.Val)
		if node.Val == maxVal {
			// nodes = append(nodes, node.Val)
			res++
		}

		dfs(node.Left, maxVal)
		dfs(node.Right, maxVal)
	}

	dfs(root, math.MinInt)
	// fmt.Println(nodes)

	return res
}
