package LinkedList

// https://leetcode.cn/problems/reverse-linked-list/

func reverseList(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
