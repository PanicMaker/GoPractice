package LinkedList

import "math"

// https://leetcode.cn/problems/add-two-numbers

// 遇到超出Int类型长度的测试用例无法通过
func addTwoNumbersI(l1 *ListNode, l2 *ListNode) *ListNode {
	listToNum := func(p *ListNode) int {
		loop, sum := 0, 0
		for p != nil {
			d := math.Pow(10, float64(loop))
			sum += p.Val * int(d)
			p = p.Next
			loop++
		}

		return sum
	}

	num1 := listToNum(l1)
	num2 := listToNum(l2)

	sum := num1 + num2

	res := &ListNode{}
	p := res

	if sum == 0 {
		return &ListNode{Val: 0}
	}

	for sum > 0 {
		d := sum % 10
		p.Next = &ListNode{Val: d}
		p = p.Next
		sum = sum / 10
	}

	return res.Next
}

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
