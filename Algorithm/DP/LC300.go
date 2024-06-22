package DP

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {
	// 创建一个长度为 nums 数组长度的动态规划数组 dp
	dp := make([]int, len(nums))

	// 初始化 dp 数组，每个元素至少可以作为一个长度为1的递增子序列的结尾
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	// 用于存储最终结果，即最长递增子序列的长度
	result := 0

	for i := 0; i < len(nums); i++ {
		// 对于每个位置 i，考虑以 nums[i] 结尾的最长递增子序列长度
		for j := 0; j < i; j++ {
			// 如果 nums[i] 大于 nums[j]，说明可以将 nums[i] 接在以 nums[j] 结尾的递增子序列后面，
			// 更新以 nums[i] 结尾的最长递增子序列的长度 dp[i]
			if nums[i] > nums[j] {
				// 如果 dp[j] + 1 不大于 dp[i]，则保持 dp[i] 不变，因为没有找到更长的递增子序列。
				dp[i] = max(dp[i], dp[j]+1)
			}
			// fmt.Printf("i = %d, j = %d, dp =%v\n", i, j, dp)
		}
		// 更新全局的最长递增子序列长度
		result = max(result, dp[i])
	}

	return result
}
