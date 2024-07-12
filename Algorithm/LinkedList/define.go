package LinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val: arr[i],
		}
		cur = cur.Next
	}

	return head
}
