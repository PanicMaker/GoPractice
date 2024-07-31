package Tree

import "strings"

// https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another/description/

// 先找出公共祖先，再分别求出公共祖先到两个节点的路径，最后拼接在一起。会超出题目内存限制
func getDirectionsI(root *TreeNode, startValue int, destValue int) string {
	var findLCA func(node *TreeNode, pVal int, qVal int) *TreeNode
	findLCA = func(node *TreeNode, pVal int, qVal int) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == pVal || node.Val == qVal {
			return node
		}

		left := findLCA(node.Left, pVal, qVal)
		right := findLCA(node.Right, pVal, qVal)

		if left != nil && right != nil {
			return node
		}

		if left != nil {
			return left
		}

		return right
	}

	lca := findLCA(root, startValue, destValue)

	var findPath func(node *TreeNode, target int, path string) string
	findPath = func(node *TreeNode, target int, path string) string {
		if node == nil {
			return ""
		}

		if node.Val == target {
			return path
		}

		left := findPath(node.Left, target, path+"L")
		right := findPath(node.Right, target, path+"R")

		return left + right
	}

	toS := findPath(lca, startValue, "")
	toS = strings.Repeat("U", len(toS))

	toD := findPath(lca, destValue, "")

	return toS + toD
}

func getDirectionsII(root *TreeNode, startValue int, destValue int) string {
	path := make([]byte, 0)

	// 辅助函数，用于查找从当前节点到目标值的路径
	var findPath func(node *TreeNode, target int) bool
	findPath = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}

		if node.Val == target {
			return true
		}

		// 尝试左子树
		path = append(path, 'L')
		if findPath(node.Left, target) {
			return true
		}

		// 如果在左子树中未找到，则尝试右子树
		path[len(path)-1] = 'R'
		if findPath(node.Right, target) {
			return true
		}

		// 如果左右子树都未找到，回溯
		path = path[:len(path)-1]
		return false
	}

	// 查找从根节点到startValue的路径
	findPath(root, startValue)
	pathToStart := path

	// 重置路径并查找从根节点到destValue的路径
	path = nil
	findPath(root, destValue)
	pathToDest := path

	// 去掉公共前缀部分
	for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
		pathToStart = pathToStart[1:] // 去掉前缀相同的部分
		pathToDest = pathToDest[1:]
	}

	// 构造最终路径: 'U'表示每次向上到最近公共祖先的步数，然后是到目标节点的剩余路径
	return strings.Repeat("U", len(pathToStart)) + string(pathToDest)
}
