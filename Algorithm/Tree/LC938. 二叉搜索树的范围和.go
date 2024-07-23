package Tree

// https://leetcode.cn/problems/range-sum-of-bst/description/

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return sum
}
