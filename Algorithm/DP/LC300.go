package DP

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	result := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}
