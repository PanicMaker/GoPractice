package LinkedList

// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/description/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
