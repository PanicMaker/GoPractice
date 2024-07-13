package LinkedList

// https://leetcode.cn/problems/rotate-list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}

	// Step 1: 计算链表长度
	p := head
	n := 0
	for p != nil {
		n++
		p = p.Next
	}

	// Step 2: 计算实际需要旋转的步数
	steps := k % n

	// Step 3: 找到新的头节点并断开链表
	// 使用虚拟节点（dummy）简化处理边界情况
	dummy := &ListNode{Next: head}

	// 快指针先移动 steps 步
	fast := head
	for i := 0; i < steps; i++ {
		fast = fast.Next
	}

	// 慢指针从 head 开始移动，直到快指针到达链表末尾
	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 此时 slow.Next 是新的头节点，需要调整指针连接
	fast.Next = head       // 将链表的末尾连接到原始头节点
	dummy.Next = slow.Next // 新的头节点是 slow.Next
	slow.Next = nil        // 断开链表连接

	// Step 4: 返回旋转后的新头节点
	return dummy.Next
}
