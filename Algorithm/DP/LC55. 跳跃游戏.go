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

// 通过贪心算法，时间复杂度为𝑂(𝑛)
func canJumpIV(nums []int) bool {
	n := len(nums)
	// 用于记录能够到达的最远位置
	farthest := 0

	for i := 0; i < n; i++ {
		// 如果当前索引 i 超过了 farthest，则返回 false，表示不能到达位置 i
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+nums[i])
	}

	return true
}
