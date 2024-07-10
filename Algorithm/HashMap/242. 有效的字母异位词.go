package HashMap

// https://leetcode.cn/problems/valid-anagram

func isAnagramI(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)

	for _, v := range s {
		mapS[v]++
	}

	for _, v := range t {
		if _, ok := mapS[v]; ok {
			mapS[v]--
			if mapS[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func isAnagramII(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
