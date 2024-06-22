package DP

// https://leetcode.cn/problems/word-break/

// 循环求解，部分用例通不过
func wordBreak1(s string, wordDict []string) bool {
	target := s

	for len(target) > 0 {
		matched := false
		for i := 0; i < len(wordDict); i++ {
			tmp := wordDict[i]
			if len(target) < len(tmp) {
				continue
			}
			if target[:len(tmp)] == tmp {
				target = target[len(tmp):]
				matched = true
				break
			}
		}

		if !matched {
			break
		}
	}
	if len(target) == 0 {
		return true
	}
	return false
}

func wordBreak2(s string, wordDict []string) bool {
	length := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}

	dp := make([]bool, length+1)
	dp[0] = true

	for i := 0; i < length; i++ {
		for j := i + 1; j < length+1; j++ {
			word := s[i:j]
			if dp[i] && wordMap[word] {
				dp[j] = true
			}
		}
	}

	return dp[length]
}
