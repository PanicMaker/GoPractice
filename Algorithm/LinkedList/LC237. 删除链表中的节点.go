package LinkedList

// https://leetcode.cn/problems/delete-node-in-a-linked-list/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
