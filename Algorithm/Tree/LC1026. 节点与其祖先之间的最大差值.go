package Tree

// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/

func maxAncestorDiff(root *TreeNode) int {

	var dfs func(node *TreeNode, minVal int, maxVal int) int
	dfs = func(node *TreeNode, minVal int, maxVal int) int {
		if node == nil {
			return maxVal - minVal
		}

		if node.Val < minVal {
			minVal = node.Val
		}
		if node.Val > maxVal {
			maxVal = node.Val
		}

		left := dfs(node.Left, minVal, maxVal)
		right := dfs(node.Right, minVal, maxVal)

		if left > right {
			return left
		}
		return right
	}

	return dfs(root, root.Val, root.Val)
}
