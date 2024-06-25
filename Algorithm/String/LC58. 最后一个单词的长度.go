package String

import (
	"regexp"
	"strings"
)

//https://leetcode.cn/problems/length-of-last-word

// 无法分割 "   fly me   to   the moon  "
func lengthOfLastWordI(s string) int {
	strs := strings.Split(s, " ")
	return len(strs[len(strs)-1])
}

// 用正则表达式
func lengthOfLastWordII(s string) int {
	reg := regexp.MustCompile(`[a-zA-Z]+`)
	strs := reg.FindAllString(s, -1)

	return len(strs[len(strs)-1])
}

// 用指针找到一个第一个不为空的字符，并开始计算长度
func lengthOfLastWordIII(s string) int {
	n := len(s)
	end := n - 1

	for s[end] == ' ' {
		end--
	}

	ans := 0
	for s[end] != ' ' {
		ans++
		end--
		if end < 0 {
			break
		}
	}

	return ans
}
