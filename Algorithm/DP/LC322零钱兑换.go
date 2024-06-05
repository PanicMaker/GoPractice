package DP

import "slices"

// https://leetcode.cn/problems/coin-change/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	slices.Sort(coins)

	n := len(coins)

	dp := make([]int, n)

	target := amount
	for i := n - 1; i >= 0; i-- {
		for j := 0; j*coins[i] <= target; j++ {
			dp[i] = j
		}
		target = target - dp[i]*coins[i]
	}

	if target == 0 {
		sum := 0
		for _, v := range dp {
			sum += v
		}
		return sum
	} else {
		return -1
	}
}
