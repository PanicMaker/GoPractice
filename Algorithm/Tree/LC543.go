package Tree

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 左子树到根节点的最长路径就等于左子树的最大深度
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)

		res = max(res, left+right)
		return max(left, right) + 1
	}

	maxDepth(root)
	return res
}
