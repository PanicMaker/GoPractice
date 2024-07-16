package LinkedList

// https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/

func modifiedList(nums []int, head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy

	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}

	for cur.Next != nil {
		if _, ok := dict[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
