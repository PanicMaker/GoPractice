package LinkedList

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/description/

func partition(head *ListNode, x int) *ListNode {
	tmp1, tmp2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	p1, p2 := tmp1, tmp2
	
	p := head
	
	for p != nil{
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		p, p.Next = p.Next, nil
	}
	
	p1.Next = tmp2.Next
	return tmp1.Next
}