package DP

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			for x := i; x < m; x++ {
				dp[x][0] = 0
			}
			break
		} else {
			dp[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			for y := j; y < n; y++ {
				dp[0][y] = 0
			}
			break
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]

			}
		}
	}

	return dp[m-1][n-1]
}
