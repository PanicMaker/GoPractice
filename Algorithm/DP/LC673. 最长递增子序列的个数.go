package DP

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence

// 无法通过 [2,2,2,2,2] 这个输入
func findNumberOfLISI(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	highest := 0
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}

			if dp[i] > highest {
				highest = dp[i]
				result = 0
				fmt.Printf("highest update to %d\n", highest)
			}

			if dp[j]+1 == highest {
				result++
			}
			fmt.Printf("i = %d, j = %d, highest = %d, result = %d, dp =%v\n", i, j, highest, result, dp)
		}
	}

	return result
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, n)
	// count[i] 表示以 nums[i] 结尾且长度为 dp[i] 的递增子序列的个数
	count := make([]int, n)

	// 初始情况下，每个元素自成一个长度为 1 的递增子序列
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}

	// maxLength 记录全局最长递增子序列的长度
	maxLength := 1
	// result 记录全局最长递增子序列的个数
	result := 0

	// 从第二个元素开始遍历数组
	for i := 0; i < n; i++ {
		// 遍历当前元素之前的所有元素
		for j := 0; j < i; j++ {
			// 只有在 nums[j] < nums[i] 的情况下才考虑更新
			if nums[j] < nums[i] {
				// 如果 dp[j] + 1 > dp[i]，则说明找到了一个更长的递增子序列，
				// 更新 dp[i] = dp[j] + 1，同时 count[i] 等于 count[j]，
				// 因为以 nums[j] 结尾的子序列都可以延长到以 nums[i] 结尾。
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
					// 如果 dp[j] + 1 == dp[i]，则表示找到了另一个长度相同的递增子序列，
					// 累加 count[j] 到 count[i]。
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}

		// 更新全局最长递增子序列的长度和对应的个数
		if dp[i] > maxLength {
			maxLength = dp[i]
			result = count[i]
		} else if dp[i] == maxLength {
			result += count[i]
		}
	}

	// 返回最终结果
	return result
}
