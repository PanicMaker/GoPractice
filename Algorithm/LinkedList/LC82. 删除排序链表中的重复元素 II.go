package LinkedList

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii

func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 创建一个虚拟头节点，指向链表的实际头节点
	dummy := &ListNode{Next: head}
	// 慢指针指向虚拟头节点
	slow := dummy

	// 遍历链表，直到链表结束
	for slow.Next != nil {
		// 快指针指向慢指针的下一个节点
		fast := slow.Next
		// 移动快指针，跳过所有重复的节点
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		// 如果慢指针的下一个节点和快指针相同，说明没有重复节点
		if slow.Next == fast {
			// 移动慢指针到下一个节点
			slow = slow.Next
		} else {
			// 否则，跳过所有重复节点
			slow.Next = fast.Next
		}
	}

	// 返回新链表的头节点
	return dummy.Next
}
