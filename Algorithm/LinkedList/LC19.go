package LinkedList

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp.Next = head

	p1 := tmp
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}

	p2 := tmp
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	// 删除节点
	p2.Next = p2.Next.Next

	return tmp.Next
}
