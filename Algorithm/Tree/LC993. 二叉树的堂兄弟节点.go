package Tree

// https://leetcode.cn/problems/cousins-in-binary-tree/description/

// 通过层次遍历（广度优先搜索，BFS）遍历二叉树，并在遍历过程中检查两个节点是否在同一层且父节点不同
func isCousinsI(root *TreeNode, x int, y int) bool {
	// 创建一个队列用于层次遍历，初始时将根节点放入队列中
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// 开始层次遍历
	for len(queue) != 0 {
		// 当前层的节点数量
		levelSize := len(queue)
		// 用于记录当前层节点的值，初始化为空
		dict := make(map[int]int)

		// 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 取出队列的第一个节点
			cur := queue[0]
			queue = queue[1:]

			// 检查当前节点是否有左右子节点，且子节点的值是否分别为x和y
			if cur.Left != nil && cur.Right != nil {
				if cur.Left.Val == x && cur.Right.Val == y {
					return false // 同一父节点的子节点，返回false
				}
				if cur.Left.Val == y && cur.Right.Val == x {
					return false // 同一父节点的子节点，返回false
				}
			}

			// 记录当前节点的值
			dict[cur.Val]++

			// 将左子节点加入队列
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			// 将右子节点加入队列
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		// 检查当前层是否只找到了x或只找到了y
		if dict[x] != 0 && dict[y] == 0 {
			return false // 只有x，没有y，返回false
		}

		if dict[x] == 0 && dict[y] != 0 {
			return false // 只有y，没有x，返回false
		}

		// 如果当前层既有x又有y，则继续下一层的遍历
	}

	// 如果遍历结束后，未能提前返回false，则说明x和y是堂兄弟节点
	return true
}

// 通过深度优先遍历
func isCousinsII(root *TreeNode, x int, y int) bool {
	depth := 0
	father := &TreeNode{Val: 0}

	var dfs func(node *TreeNode, fa *TreeNode, d int) bool
	dfs = func(node *TreeNode, fa *TreeNode, d int) bool {
		if node == nil {
			return false
		}

		if node.Val == x || node.Val == y {
			if depth > 0 {
				return depth == d && father != fa
			}
			depth, father = d, fa
		}

		return dfs(node.Left, node, d+1) || dfs(node.Right, node, d+1)
	}

	return dfs(root, nil, 1)
}
