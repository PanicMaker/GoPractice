package Tree

// https://leetcode.cn/problems/cousins-in-binary-tree-ii/description/

func replaceValueInTree(root *TreeNode) *TreeNode {
	sums := make(map[int]int)

	var calcSum func(node *TreeNode, depth int)
	calcSum = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		sums[depth] += node.Val

		calcSum(node.Left, depth+1)
		calcSum(node.Right, depth+1)
	}

	calcSum(root, 0)

	var update func(node *TreeNode, depth int)
	update = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Left != nil && node.Right != nil {
			res := sums[depth+1] - node.Left.Val - node.Right.Val
			node.Left.Val = res
			node.Right.Val = res
		} else {
			if node.Left != nil {
				node.Left.Val = sums[depth+1] - node.Left.Val
			}

			if node.Right != nil {
				node.Right.Val = sums[depth+1] - node.Right.Val
			}
		}

		update(node.Left, depth+1)
		update(node.Right, depth+1)
	}

	update(root, 0)
	root.Val = 0

	return root
}
