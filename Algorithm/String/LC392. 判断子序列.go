package String

import (
	"strings"
)

// https://leetcode.cn/problems/is-subsequence/

// 使用动态规划解答，时间复杂度是 O(n*m)
func isSubsequenceI(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	if n > len(t) {
		return false
	}

	dp := make([]bool, n)
	var index int
	if index = strings.Index(t, string(s[0])); index >= 0 {
		dp[0] = true
	} else {
		dp[0] = false
		return false
	}

	j := index + 1
	for i := 1; i < n; i++ {
		for j < len(t) {
			if dp[i-1] && s[i] == t[j] {
				dp[i] = true
				j++
				break
			}
			j++
		}
	}

	return dp[n-1]
}

// 使用动态规划优化版，时间复杂度是 O(n*m)
func isSubsequenceII(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	dp := make([]bool, n+1)
	// dp[0]代表空字符串""
	dp[0] = true

	for i := 1; i <= n; i++ {
		index := strings.Index(t, string(s[i-1]))
		if index >= 0 && dp[i-1] {
			dp[i] = true
			t = t[index+1:] // 更新 t，从找到的位置后一位开始继续查找
		} else {
			dp[i] = false
		}
	}

	return dp[n-1]
}

// 用双指针同时对两个字符串进行比较
func isSubsequenceIII(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 {
		return true
	}

	i, j := 0, 0
	for j < m {
		if s[i] == t[j] {
			i++
			if i == n {
				return true
			}
		}
		j++
	}

	return false
}
