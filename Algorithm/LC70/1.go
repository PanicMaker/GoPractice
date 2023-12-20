package main

import "fmt"

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs2(n int) int {
	prev := 1
	cur := 2
	for i := 3; i <= n; i++ {
		temp := cur
		cur = prev + cur
		prev = temp
	}
	return cur
}

func main() {
	fmt.Println(climbStairs2(3))
}
