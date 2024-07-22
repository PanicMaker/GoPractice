package Tree

// https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/

func sufficientSubset(root *TreeNode, limit int) *TreeNode {

	var dfs func(node *TreeNode, sum int) *TreeNode
	dfs = func(node *TreeNode, sum int) *TreeNode {
		if node == nil {
			return nil
		}

		sum += node.Val

		if node.Left == nil && node.Right == nil {
			if sum < limit {
				return nil
			}
			return node
		}

		node.Left = dfs(node.Left, sum)
		node.Right = dfs(node.Right, sum)

		// 如果左子树和右子树都被删除了
		if node.Left == nil && node.Right == nil {
			return nil
		}

		return node
	}

	return dfs(root, 0)
}
