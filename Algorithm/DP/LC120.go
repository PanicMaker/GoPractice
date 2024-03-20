package DP

// 120. 三角形最小路径和

func minimumTotal(triangle [][]int) int {
	m := len(triangle)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	lastrow := triangle[m-1]

	for i := 0; i < m; i++ {
		dp[m-1][i] = lastrow[i]
	}

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}
