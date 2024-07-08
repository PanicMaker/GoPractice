package stack

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		switch v {
		case "+":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d1 + d2
			stack = append(stack, res)
		case "-":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d2 - d1
			stack = append(stack, res)
		case "*":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d1 * d2
			stack = append(stack, res)
		case "/":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d2 / d1
			stack = append(stack, res)
		default:
			d, _ := strconv.Atoi(v)
			stack = append(stack, d)
		}
	}

	return stack[len(stack)-1]
}
