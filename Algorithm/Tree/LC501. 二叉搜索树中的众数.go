package Tree

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	dict := make(map[int]int)
	maxNum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dict[node.Val]++
		maxNum = max(maxNum, dict[node.Val])

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	for k, v := range dict {
		if v == maxNum {
			res = append(res, k)
		}
	}

	return res
}
