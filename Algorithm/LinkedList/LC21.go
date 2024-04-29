package LinkedList

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmp := &ListNode{-1, nil}
	p := tmp
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return tmp.Next
}
