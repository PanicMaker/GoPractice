package Tree

// https://leetcode.cn/problems/delete-nodes-and-return-forest/

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res := make([]*TreeNode, 0)

	dict := make(map[int]int)
	for _, v := range to_delete {
		dict[v]++
	}

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if _, ok := dict[node.Val]; ok {
			if node.Left != nil {
				res = append(res, node.Left)
			}
			if node.Right != nil {
				res = append(res, node.Right)
			}
			node = nil
		}
		return node
	}

	root = dfs(root)
	if root != nil {
		res = append(res, root)
	}

	return res
}
