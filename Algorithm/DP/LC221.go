package DP

// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square/description

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	maxSide := 0

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			// 如果在矩阵的边上，则存在一个边至少为1的正方形
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
			maxSide = max(dp[i][j], maxSide)
		}
	}

	return maxSide * maxSide
}
