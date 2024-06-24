package DP

// https://leetcode.cn/problems/jump-game/description

// 题目中超长切片测试用例会超时
func canJumpI(nums []int) bool {
	n := len(nums)

	dp := make([]bool, n)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for step := 0; step <= nums[j]; step++ {
				if dp[i] {
					break
				}
				dp[i] = dp[j] && step+j == i
			}
		}
	}

	return dp[n-1]
}

// 时间复杂度O(n²)
func canJumpII(nums []int) bool {
	n := len(nums)

	dp := make([]bool, n)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if i-j <= nums[j] && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[n-1]
}

func canJumpIII(nums []int) bool {
	jumpHighest := 0

	for i := 0; i <= jumpHighest && jumpHighest < len(nums); i++ {
		if jumpHighest < i+nums[i] {
			jumpHighest = i + nums[i]
		}
	}

	return jumpHighest >= len(nums)-1
}
