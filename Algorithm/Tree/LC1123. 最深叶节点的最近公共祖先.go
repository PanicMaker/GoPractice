package Tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/description/

func lcaDeepestLeavesI(root *TreeNode) *TreeNode {
	// 1. 找到最深的节点
	queue := []*TreeNode{root}
	var currentLevel []*TreeNode

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel = []*TreeNode{}

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, current)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	// 2. 找到最深节点的最近公共祖先
	var dfs func(root *TreeNode, nodes []*TreeNode) *TreeNode
	dfs = func(root *TreeNode, nodes []*TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		for _, node := range nodes {
			if root == node {
				return root
			}
		}

		left := dfs(root.Left, nodes)
		right := dfs(root.Right, nodes)

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	return dfs(root, currentLevel)
}

func lcaDeepestLeavesII(root *TreeNode) *TreeNode {
	maxDepth := -1

	var ans *TreeNode

	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			maxDepth = max(maxDepth, depth)
			return depth
		}

		left := dfs(node.Left, depth+1)
		right := dfs(node.Right, depth+1)

		if left == right && left == maxDepth {
			ans = node
		}

		return max(left, right)
	}

	dfs(root, 0)

	return ans
}
