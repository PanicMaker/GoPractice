package DP

// 64. 最小路径和
// https://leetcode.cn/problems/minimum-path-sum/description/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 设置 dp 的边界
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
