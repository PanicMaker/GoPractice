package String

import "strings"

// https://leetcode.cn/problems/longest-common-prefix

func longestCommonPrefixI(strs []string) string {
	// 选定数组第一个字符串去匹配
	match := strs[0]

	for _, str := range strs {
		for !strings.HasPrefix(str, match) {
			if len(match) == 0 {
				return ""
			}
			// 若不匹配，截去字符串最后一位重新匹配
			match = match[:len(match)-1]
		}
	}

	return match
}

func longestCommonPrefixII(strs []string) string {
	match := strs[0]

	// 循环数组第一个字符串的每一列
	for i := 0; i < len(strs[0]); i++ {
		// 循环数组中除第一个字符串的剩余字符串
		for j := 1; j < len(strs); j++ {
			// 将字符串的每一列进行对比
			if i == len(strs[j]) || match[i] != strs[j][i] {
				return match[:i]
			}
		}
	}

	return match
}
