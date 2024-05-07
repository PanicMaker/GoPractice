package LinkedList

// 876. 链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/description/

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	
	return slow
}