package HashMap

import (
	"strings"
)

// https://leetcode.cn/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")

	if len(pattern) != len(strs) {
		return false
	}

	mapPS := make(map[byte]string)
	mapSP := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		if v, ok := mapPS[pattern[i]]; ok {
			if v != strs[i] {
				return false
			}
		} else {
			mapPS[pattern[i]] = strs[i]
		}

		if v, ok := mapSP[strs[i]]; ok {
			if v != pattern[i] {
				return false
			}
		} else {
			mapSP[strs[i]] = pattern[i]
		}
	}

	return true
}
