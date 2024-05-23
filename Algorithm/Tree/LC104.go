package Tree

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

// 通过二叉树遍历的方式求解
func maxDepth1(root *TreeNode) int {
	var res, depth int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		depth++

		traverse(node.Left)
		traverse(node.Right)
		res = max(res, depth)
		// if node.Left == nil && node.Right == nil {
		// 	res = max(res, depth)
		// }
		// depth--
	}

	traverse(root)

	return res
}

// 分别求出左右子树的最大深度
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res1 := maxDepth2(root.Left)
	res2 := maxDepth2(root.Right)
	res := max(res1, res2) + 1
	return res
}
