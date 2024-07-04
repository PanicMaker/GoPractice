package String

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	need := make(map[string]int)
	window := make(map[string]int)

	for _, v := range words {
		need[v]++
	}

	res := make([]int, 0)

	n := len(words)
	wordLength := len(words[0])

	left, right := 0, wordLength
	valid := 0

	if len(s) < n*wordLength {
		return res
	}

	for right < len(s) {
		str := s[right-wordLength : right]

		if _, ok := need[str]; ok {
			right += wordLength
			window[str]++
			if window[str] == need[str] {
				valid++
			}
		} else {
			right++
		}

		for right-left >= n*wordLength && left < len(s)-n*wordLength {
			if valid == len(words) {
				res = append(res, left)
			}

			d := s[left : left+wordLength]
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}
