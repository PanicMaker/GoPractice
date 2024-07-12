package String

import (
	"strings"
)

// https://leetcode.cn/problems/reverse-words-in-a-string

func reverseWords(s string) string {
	// 去除字符串开头的空字符串
	s = strings.TrimSpace(s)
	n := len(s)
	start, end := 0, 0
	words := make([]string, 0)

	for start < n {
		// for end < n && s[end] == ' '{
		// 	end++
		// }
		// start = end

		for end < n && s[end] != ' ' {
			end++
		}

		words = append(words, s[start:end])

		for end < n && s[end] == ' ' {
			end++
		}
		start = end
	}

	left, right := 0, len(words)-1

	// 将数组翻转
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	res := strings.Join(words, " ")

	return res
}
