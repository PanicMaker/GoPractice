package Tree

// https://leetcode.cn/problems/binary-tree-right-side-view/

// 深度优先遍历解题
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, node.Val)
		}

		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)

	return res
}

// 广度优先遍历解题
func rightSideViewII(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		res = append(res, queue[n-1].Val)

		newQueue := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				newQueue = append(newQueue, v.Left)
			}
			if v.Right != nil {
				newQueue = append(newQueue, v.Right)
			}
		}
		queue = newQueue
	}

	return res
}
