package DP

import "math"

// 265. 粉刷房子II
// https://leetcode.cn/problems/paint-house-ii/

func minCostII(costs [][]int) int {
	m, n := len(costs), len(costs[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := math.MaxInt
			for x := 0; x < n; x++ {
				if x != j {
					// 找到在i-1行中dp的最小值
					temp = min(dp[i-1][x], temp)
					dp[i][j] = temp + costs[i][j]
				}
			}
		}
	}

	result := math.MaxInt
	for i := 0; i < n; i++ {
		result = min(result, dp[m-1][i])
	}

	return result
}
