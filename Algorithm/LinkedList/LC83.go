package LinkedList

// 83. 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1 := head

	for p2 := head; p2 != nil; p2 = p2.Next {
		if p1.Val != p2.Val {
			p1.Next = p2
			p1 = p1.Next
		}
	}
	p1.Next = nil
	return head
}
