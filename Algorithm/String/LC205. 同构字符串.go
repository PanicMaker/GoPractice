package String

// https://leetcode.cn/problems/isomorphic-strings/description/

import (
	"reflect"
	"strings"
)

func isIsomorphicI(s string, t string) bool {
	StrToSlice := func(s string) []int {
		arr := make([]int, 0)
		for _, v := range s {
			arr = append(arr, strings.Index(s, string(v)))
		}
		return arr
	}

	arrS := StrToSlice(s)
	arrT := StrToSlice(t)

	if reflect.DeepEqual(arrS, arrT) {
		return true
	}
	return false
}

func isIsomorphicII(s string, t string) bool {
	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {

		if v, ok := mapS[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			mapS[s[i]] = t[i]
		}

		if v, ok := mapT[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			mapT[t[i]] = s[i]
		}
	}

	return true
}
