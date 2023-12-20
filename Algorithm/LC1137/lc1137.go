package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
}

func main() {
	fmt.Println(tribonacci(34))
}
