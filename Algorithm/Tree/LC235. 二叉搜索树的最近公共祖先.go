package Tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x {
		return lowestCommonAncestorBST(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	return root
}
