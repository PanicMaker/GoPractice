package Tree

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

var res, depth int

func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	depth++
	traverse(node.Left)
	traverse(node.Right)
	if node.Left == nil && node.Right == nil {
		res = max(res, depth)
	}
	depth--
}
