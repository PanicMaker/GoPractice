package String

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 动态规划解法
func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	maxLength := 1 // 最长回文子串的长度，至少为1
	begin := 0     // 最长回文子串的起始位置

	// 初始化动态规划数组
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		// 单个字符一定是回文串
		dp[i][i] = true
	}

	// 从后往前遍历字符串
	for i := n - 1; i >= 0; i-- {
		// 从当前字符到字符串末尾遍历
		for j := i; j < n; j++ {
			if i == j {
				// 单个字符一定是回文串
				dp[i][j] = true
			} else if j == i+1 {
				// 两个相邻字符相同则是回文串
				dp[i][j] = s[i] == s[j]
			} else if s[i] == s[j] && dp[i+1][j-1] {
				// 当前字符与末尾字符相同，并且去掉头尾后的子串也是回文串，则当前子串是回文串
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > maxLength {
				// 如果当前子串是回文串，并且长度大于之前记录的最大长度，则更新最大长度和起始位置
				maxLength = j - i + 1
				begin = i
			}
		}
	}

	// 返回最长回文子串
	return s[begin : begin+maxLength]
}

// 使用双指针解决
func longestPalindrome2(s string) string {
	// 在 s 中寻找以 s[l] 和 s[r] 为中心的最长回文串
	palindrome := func(str string, l int, r int) string {
		for l >= 0 && r < len(str) && str[l] == str[r] {
			l--
			r++
		}
		return str[l+1 : r]
	}

	res := ""
	for i := 0; i < len(s); i++ {
		// 奇数回文串
		s1 := palindrome(s, i, i)
		// 偶数回文串
		s2 := palindrome(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}

	return res
}
