package main

// https://leetcode.cn/problems/fibonacci-number/description

func fib(n int) int {
	prev := 0
	cur := 1

	for i := 1; i < n; i++ {
		temp := cur
		cur = prev + cur
		prev = temp
	}

	return cur
}
