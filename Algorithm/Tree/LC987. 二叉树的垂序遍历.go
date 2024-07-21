package Tree

import (
	"sort"
)

// https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/description/

func verticalTraversal(root *TreeNode) [][]int {
	type KeyValue struct {
		Val int
		Row int
	}

	res := make([][]int, 0)
	dict := make(map[int][]KeyValue)

	var dfs func(node *TreeNode, row int, col int)
	dfs = func(node *TreeNode, row int, col int) {
		if node == nil {
			return
		}

		// 将节点的值和行号存入对应的列中
		dict[col] = append(dict[col], KeyValue{Val: node.Val, Row: row})

		// 递归遍历左右子树，列号调整
		dfs(node.Left, col-1, row+1)
		dfs(node.Right, col+1, row+1)
	}

	// 从根节点开始深度优先遍历
	dfs(root, 0, 0)

	// 提取列号并排序
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 构建最终结果
	for _, col := range keys {
		// 提取当前列的所有节点数据
		values := dict[col]

		// 按照行号对节点数据进行排序
		sort.Slice(values, func(i, j int) bool {
			return values[i].Row < values[j].Row || (values[i].Row == values[j].Row && values[i].Val < values[j].Val)
		})

		// 只保留节点的值
		column := make([]int, len(values))
		for i, kv := range values {
			column[i] = kv.Val
		}

		// 将当前列的结果添加到最终结果中
		res = append(res, column)
	}

	return res
}
