package LinkedList

// https://leetcode.cn/problems/reverse-linked-list-ii

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	p := dummy
	// 将p指向left节点的前一节点
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	var pre *ListNode
	cur := p.Next // 指向值为left的节点

	// 将从left到right的节点进行翻转
	for i := 0; i < right-left+1; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// 将right节点之后的链表接入left节点之后
	p.Next.Next = cur
	// 接上翻转后的链表
	p.Next = pre

	return dummy.Next
}
