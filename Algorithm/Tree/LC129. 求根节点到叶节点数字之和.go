package Tree

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/

func sumNumbers(root *TreeNode) int {
	nums := make([]int, 0)
	sum := 0

	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}

		num = num*10 + node.Val

		if node.Left == nil && node.Right == nil {
			nums = append(nums, num)
		}

		dfs(node.Left, num)
		dfs(node.Right, num)
	}

	dfs(root, 0)

	for _, v := range nums {
		sum += v
	}

	return sum
}
