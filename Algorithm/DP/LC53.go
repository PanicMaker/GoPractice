package DP

// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/description/

func maxSubArray(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	dp[0] = nums[0]

	result := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}

	return result
}
