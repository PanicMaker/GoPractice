package LinkedList

// https://leetcode.cn/problems/palindrome-linked-list/

func isPalindrome(head *ListNode) bool {
	// 找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转中点后的链表
	cur, pre := slow, (*ListNode)(nil)
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// 循环比较两端链表的节点值
	p1, p2 := head, pre
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

func isPalindromeII(head *ListNode) bool {
	left := head

	is := true

	var traverse func(right *ListNode)
	traverse = func(right *ListNode) {
		if right == nil {
			return
		}

		// 利用递归，走到链表尾部
		traverse(right.Next)

		// 后序遍历位置，此时的right指针指向链表右侧尾部
		// 所以可以和left指针比较，判断是否是回文链表
		if left.Val != right.Val {
			is = false
		}

		left = left.Next
	}

	traverse(head)

	return is
}
