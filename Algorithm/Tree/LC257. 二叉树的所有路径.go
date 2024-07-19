package Tree

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/binary-tree-paths/

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if path == "" {
			path = strconv.Itoa(node.Val)
		} else {
			path = fmt.Sprintf("%s->%d", path, node.Val)
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		}

		dfs(node.Left, path)
		dfs(node.Right, path)
	}

	dfs(root, "")

	return res
}
