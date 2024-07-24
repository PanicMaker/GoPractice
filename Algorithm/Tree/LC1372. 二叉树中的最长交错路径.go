package Tree

// https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/description/

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	var dfs func(node *TreeNode, direction int, length int)
	dfs = func(node *TreeNode, direction int, length int) {
		if node == nil {
			return
		}

		ans = max(ans, length)

		// 根据当前方向决定下一步的方向
		if direction == -1 {
			dfs(node.Left, 1, length+1) // 左变右
			dfs(node.Right, -1, 1)      // 左变左，从右子树开始新路
		} else if direction == 1 {
			dfs(node.Right, -1, length+1) // 右变左
			dfs(node.Left, 1, 1)          // 右变右，从左子树开始新路径
		}
	}

	dfs(root, -1, 0)
	dfs(root, 1, 0)

	return ans
}
