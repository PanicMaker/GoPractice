package DP

import "math"

// https://leetcode.cn/problems/jump-game-ii/

// 使用动态规划解决，时间复杂度为O(n²)
func jumpI(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

// 使用贪心算法，时间复杂度 O(n)
func jumpII(nums []int) int {
	n := len(nums)

	count := 0
	// 当前跳跃结束的位置
	current := 0
	// 能够到达最远的位置
	farthest := 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == current {
			count++
			current = farthest
			if current >= n-1 {
				break
			}
		}
	}

	return count
}
