package Tree

import "slices"

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		tmp := make([]int, 0)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, tmp)
	}

	// 将结果中排序奇数的数组翻转
	for i := 0; i < len(ans); i++ {
		if i%2 == 1 {
			slices.Reverse(ans[i])
		}
	}

	return ans
}
