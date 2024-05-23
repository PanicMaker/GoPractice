package DP

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description/

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
	if n == 1 {
		return 1
	}

	prev := 1
	cur := 2
	for i := 3; i <= n; i++ {
		cur, prev = prev+cur, cur
	}
	return cur
}

// 用递归的方式解决
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs3(n-1) + climbStairs3(n-2)
}
