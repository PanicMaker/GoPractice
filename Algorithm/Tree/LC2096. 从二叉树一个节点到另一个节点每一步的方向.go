package Tree

import "strings"

// https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another/description/

// 先找出公共祖先，再分别求出公共祖先到两个节点的路径，最后拼接在一起。会超出题目内存限制
func getDirectionsI(root *TreeNode, startValue int, destValue int) string {
	var dfs func(node *TreeNode, pVal int, qVal int) *TreeNode
	dfs = func(node *TreeNode, pVal int, qVal int) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == pVal || node.Val == qVal {
			return node
		}

		left := dfs(node.Left, pVal, qVal)
		right := dfs(node.Right, pVal, qVal)

		if left != nil && right != nil {
			return node
		}

		if left != nil {
			return left
		}

		return right
	}

	lca := dfs(root, startValue, destValue)

	var findPath func(node *TreeNode, target int, path string) string
	findPath = func(node *TreeNode, target int, path string) string {
		if node == nil {
			return ""
		}

		if node.Val == target {
			return path
		}

		left := findPath(node.Left, target, path+"L")
		right := findPath(node.Right, target, path+"R")

		return left + right
	}

	toS := findPath(lca, startValue, "")
	toS = strings.Repeat("U", len(toS))

	toD := findPath(lca, destValue, "")

	return toS + toD
}
