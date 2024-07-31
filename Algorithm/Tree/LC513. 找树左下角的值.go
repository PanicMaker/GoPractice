package Tree

// https://leetcode.cn/problems/find-bottom-left-tree-value/description/

func findBottomLeftValueI(root *TreeNode) int {
	ans := make([][]int, 0)

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

	// 层序遍历结果中的最后一个数组的第一个节点即为最底层最左边的节点值
	return ans[len(ans)-1][0]
}

func findBottomLeftValueII(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var node *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	// 从右往左进行层序遍历，最后一个节点即为最底层最左边的节点
	return node.Val
}
