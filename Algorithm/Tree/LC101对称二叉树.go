package Tree

// https://leetcode.cn/problems/symmetric-tree/

// 递归实现
func isSymmetric1(root *TreeNode) bool {

	var check func(p *TreeNode, q *TreeNode) bool
	check = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root.Left, root.Right)
}

// 迭代实现
func isSymmetric2(root *TreeNode) bool {
	var queue []*TreeNode

	queue = append(queue, root, root)

	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		queue = queue[2:]

		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left)
		queue = append(queue, q.Right)

		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}
	return true
}
