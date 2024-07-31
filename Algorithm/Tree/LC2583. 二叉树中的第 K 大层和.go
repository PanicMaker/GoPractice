package Tree

import "sort"

// https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/description/

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	res := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, sum)
	}

	sort.Ints(res)

	if len(res) < k {
		return -1
	}

	return int64(res[len(res)-k])
}
