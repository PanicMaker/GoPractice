package DP

// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/description/

func robII(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	subRob := func(nums []int) int {
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

	return max(subRob(nums[:n-1]), subRob(nums[1:]))
}
