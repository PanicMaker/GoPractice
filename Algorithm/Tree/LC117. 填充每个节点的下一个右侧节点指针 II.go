package Tree

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/

func connect117(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i != levelSize-1 {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
