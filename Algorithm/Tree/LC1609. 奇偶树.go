package Tree

import (
	"math"
)

// https://leetcode.cn/problems/even-odd-tree/description/

func isEvenOddTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		pre := 0

		if depth%2 == 0 {
			pre = math.MinInt
		} else if depth%2 == 1 {
			pre = math.MaxInt
		}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if depth%2 == 0 {
				if node.Val%2 == 0 || node.Val <= pre {
					return false
				}
			} else {
				if node.Val%2 == 1 || node.Val >= pre {
					return false
				}
			}

			pre = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return true
}
