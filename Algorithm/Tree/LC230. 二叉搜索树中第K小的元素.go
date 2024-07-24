package Tree

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	flag := 1

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if flag == k {
			res = node.Val
		}
		flag++
		dfs(node.Right)
	}

	dfs(root)

	return res
}
