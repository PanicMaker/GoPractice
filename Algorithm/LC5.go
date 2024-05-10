package main

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 动态规划解法
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	maxLength := 1
	begin := 0

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = false
		}
	}

	for i := range dp {
		dp[i][i] = true
	}

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				dp[i][j] = true
			} else if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLength]
}
