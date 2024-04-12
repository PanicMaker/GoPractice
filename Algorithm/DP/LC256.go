package DP

// 256. 粉刷房子
// https://leetcode.cn/problems/paint-house/

func minCost(costs [][]int) int {
	m := len(costs)
	if m == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < len(costs[0]); i++ {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}

	return min(dp[m-1][0], dp[m-1][1], dp[m-1][2])
}
