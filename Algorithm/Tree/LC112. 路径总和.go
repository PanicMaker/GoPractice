package Tree

// https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	var res bool

	var dfs func(node *TreeNode, target int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		newTarget := target - node.Val
		if node.Left == nil && node.Right == nil && newTarget == 0 {
			res = true
		}
		dfs(node.Left, newTarget)
		dfs(node.Right, newTarget)
	}

	dfs(root, targetSum)

	return res
}

func hasPathSumII(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
