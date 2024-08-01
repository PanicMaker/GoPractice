package Tree

// https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/description/

// 层序遍历方式
func reverseOddLevelsI(root *TreeNode) *TreeNode {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	even := true

	for len(queue) > 0 {
		levelSize := len(queue)
		tmp := make([]*TreeNode, 0)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			tmp = append(tmp, cur)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if !even {
			for i, j := 0, len(tmp)-1; i < j; i++ {
				tmp[i].Val, tmp[j].Val = tmp[j].Val, tmp[i].Val
				j--
			}
		}

		even = !even
	}

	return root
}

func reverseOddLevelsII(root *TreeNode) *TreeNode {

	var dfs func(root1, root2 *TreeNode, isOdd bool)
	dfs = func(root1, root2 *TreeNode, isOdd bool) {
		if root1 == nil {
			return
		}

		if isOdd {
			root1.Val, root2.Val = root2.Val, root1.Val
		}

		dfs(root1.Left, root2.Right, !isOdd)
		dfs(root1.Right, root2.Left, !isOdd)
	}

	dfs(root.Left, root.Right, true)

	return root
}
