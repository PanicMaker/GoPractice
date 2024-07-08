package stack

import "math"

// https://leetcode.cn/problems/min-stack/description

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{stack: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	if len(s.stack) == 0 {
		s.min = val
	}
	s.stack = append(s.stack, val)
	if s.min > val {
		s.min = val
	}
}

func (s *MinStack) Pop() {
	d := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	if d == s.min {
		s.min = math.MaxInt
		for i := 0; i < len(s.stack); i++ {
			if s.stack[i] <= s.min {
				s.min = s.stack[i]
			}
		}
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
