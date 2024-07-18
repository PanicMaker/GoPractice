package Tree

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/

import (
	"container/list"
	"math"
)

// minDepthI 计算二叉树的最小深度。
// 返回二叉树的最小深度。
func minDepthI(root *TreeNode) int {
	// 如果根节点为空，则树的深度为 0。
	if root == nil {
		return 0
	}

	// 初始化一个变量来存储最小深度，初始值为整型的最大值。
	var res = math.MaxInt

	// 定义一个递归函数来遍历树并计算深度。
	var traverse func(node *TreeNode, depth int)
	traverse = func(node *TreeNode, depth int) {
		// 如果当前节点为空，直接返回。
		if node == nil {
			return
		}

		// 增加当前节点的深度。
		depth++

		// 递归遍历左右子树。
		traverse(node.Left, depth)
		traverse(node.Right, depth)

		// 如果当前节点是叶子节点（左右子节点均为空），
		// 更新最小深度为当前路径的深度。
		if node.Left == nil && node.Right == nil {
			res = min(res, depth)
		}

		// 回溯：处理完当前节点后，减少深度。
		depth--
	}

	// 从根节点开始遍历，初始深度为 0。
	traverse(root, 0)

	// 返回找到的最小深度。
	return res
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minDepthBFS(root *TreeNode) int {
	// 如果根节点为空，则深度为 0。
	if root == nil {
		return 0
	}

	// 使用队列来进行 BFS，同时记录节点和其对应的深度。
	queue := list.New()
	queue.PushBack(root)
	depth := 1 // 初始深度为 1

	for queue.Len() > 0 {
		levelSize := queue.Len()

		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			// 如果当前节点是叶子节点，则返回当前深度。
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// 将当前节点的子节点加入队列。
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		// 每遍历完一层节点，深度加一。
		depth++
	}

	return depth
}
