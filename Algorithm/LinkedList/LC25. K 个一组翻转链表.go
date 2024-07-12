package LinkedList

// https://leetcode.cn/problems/reverse-nodes-in-k-group

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 求出链表的节点个数
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	// 创建一个虚拟头节点并指向实际头节点
	dummy := &ListNode{Next: head}
	p := dummy

	// 循环处理每一组长度为 k 的子链表
	for n >= k {
		// 用于反转子链表的指针
		cur := p.Next
		pre := p

		// 反转当前子链表
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		// 处理反转后的子链表和原链表的连接
		tail := p.Next  // 当前子链表的头节点在反转后会成为尾节点
		tail.Next = cur // 将原子链表的尾节点连接到反转后子链表的下一个节点
		p.Next = pre    // 将原子链表头节点连接到反转后的子链表头节点
		p = tail        // 移动 p 指针到下一个待处理的子链表头节点

		n -= k // 减去已处理的节点数量
	}

	// 返回新链表的头节点
	return dummy.Next
}
