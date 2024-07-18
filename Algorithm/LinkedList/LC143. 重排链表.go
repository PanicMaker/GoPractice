package LinkedList

// https://leetcode.cn/problems/reorder-list/

func reorderList(head *ListNode) {
	// 找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	// 翻转中点后的链表
	cur, pre := mid, (*ListNode)(nil)
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	head2 := pre
	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}
