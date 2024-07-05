package String

import (
	"strings"
)

// 1422. 分割字符串的最大得分
// https://leetcode.cn/problems/maximum-score-after-splitting-a-string/description/

func maxScore(s string) int {
	length := len(s)
	res := 0

	for i := 1; i < length; i++ {
		sum := strings.Count(s[:i], "0") + strings.Count(s[i:], "1")
		res = max(res, sum)
	}

	return res
}
