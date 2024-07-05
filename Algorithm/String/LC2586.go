package String

import (
	"fmt"
	"strings"
)

// 2586. 统计范围内的元音字符串数
// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/description/

func vowelStrings(words []string, left int, right int) int {
	matchs := "aeiou"
	count := 0

	for i := left; i <= right; i++ {
		first := words[i][:1]
		last := words[i][len(words[i])-1:]
		if strings.Contains(matchs, first) && strings.Contains(matchs, last) {
			fmt.Println(words[i])
			count++
		}
	}
	return count
}
