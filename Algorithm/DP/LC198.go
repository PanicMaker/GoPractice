package DP

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/

func rob1(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		// 当前房屋要不要偷
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

func rob2(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	pre, cur := nums[0], max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		pre, cur = cur, max(cur, pre+nums[i])

	}
	return cur
}
