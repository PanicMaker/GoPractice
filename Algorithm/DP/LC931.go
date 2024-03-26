package DP

import "fmt"

// 931. 下降路径最小和
// https://leetcode.cn/problems/minimum-falling-path-sum/description/?envType=study-plan-v2&envId=dynamic-programming

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	lastrow := matrix[m-1]

	for i := 0; i < m; i++ {
		dp[m-1][i] = lastrow[i]
	}

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if j == 0 {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + matrix[i][j]
			} else if j == m-1 {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j-1], dp[i+1][j+1]) + matrix[i][j]
			}
		}
		fmt.Println(dp)
	}

	minVal := dp[0][0]

	for j := 0; j < m; j++ {
		if minVal > dp[0][j] {
			minVal = dp[0][j]
		}
	}

	return minVal
}
