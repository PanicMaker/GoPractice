package Tree

import "math"

// https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/description/

// 时间复杂度O(n²)，测试用例会超时
func maxSumBST(root *TreeNode) int {
	res := 0
	dict := make(map[*TreeNode]int)

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		sum := node.Val + left + right

		if isBST(node) {
			dict[node] = sum
		}
		return sum
	}

	dfs(root)

	for _, v := range dict {
		res = max(res, v)
	}

	return res
}

func isBST(root *TreeNode) bool {

	var check func(node *TreeNode, left int, right int) bool
	check = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}

		// 左子树节点值应该都小于当前节点值，右子树节点值应该都大于当前节点值
		return left < node.Val && node.Val < right && check(node.Left, left, node.Val) && check(node.Right, node.Val, right)
	}

	return check(root, -math.MaxInt, math.MaxInt)
}

func maxSumBSTII(root *TreeNode) int {
	res := 0
	type result struct {
		isBst    bool
		sum      int
		maxLeft  int
		minRight int
	}

	var dfs func(node *TreeNode) result
	dfs = func(node *TreeNode) result {

		if node == nil {
			return result{
				isBst:    true,
				sum:      0,
				maxLeft:  math.MinInt,
				minRight: math.MaxInt,
			}
		}

		leftRes := dfs(node.Left)
		rightRes := dfs(node.Right)

		if leftRes.isBst && rightRes.isBst && leftRes.maxLeft < node.Val && rightRes.minRight > node.Val {
			currentSum := node.Val + leftRes.sum + rightRes.sum
			res = max(res, currentSum)
			return result{
				isBst:    true,
				sum:      currentSum,
				maxLeft:  max(node.Val, rightRes.maxLeft), // 更新maxLeft和minRight
				minRight: min(node.Val, leftRes.minRight),
			}
		} else {
			return result{
				isBst:    false,
				sum:      0,
				maxLeft:  math.MinInt,
				minRight: math.MaxInt,
			}
		}
	}

	dfs(root)

	return res
}
