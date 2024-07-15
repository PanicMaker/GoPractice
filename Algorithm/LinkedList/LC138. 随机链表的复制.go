package LinkedList

// https://leetcode.cn/problems/copy-list-with-random-pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummy := &Node{}
	dict := make(map[*Node]*Node)
	q := dummy

	p := head
	for p != nil {
		q.Next = &Node{Val: p.Val}
		q = q.Next
		dict[p] = q
		p = p.Next
	}

	p = head
	q = dummy.Next
	for p != nil {
		if p.Random != nil {
			q.Random = dict[p.Random]
		}
		p = p.Next
		q = q.Next
	}

	return dummy.Next
}
