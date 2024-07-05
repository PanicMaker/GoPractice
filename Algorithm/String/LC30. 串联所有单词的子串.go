package String

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	need := make(map[string]int)

	for _, v := range words {
		need[v]++
	}

	res := make([]int, 0)

	n := len(words)
	wordLength := len(words[0])

	if len(s) < n*wordLength {
		return res
	}

	for i := 0; i < wordLength; i++ {
		window := make(map[string]int)
		left, right := i, i
		valid := 0

		for right+wordLength <= len(s) { // 保证窗口不会越界
			str := s[right : right+wordLength]
			right += wordLength

			if count, ok := need[str]; ok {
				window[str]++
				if window[str] == count {
					valid++
				}
			}

			for right-left == n*wordLength { // 窗口长度等于所有单词的总长度
				if valid == len(need) {
					res = append(res, left)
				}

				d := s[left : left+wordLength]
				left += wordLength

				if count, ok := need[d]; ok {
					if window[d] == count {
						valid--
					}
					window[d]--
				}
			}
		}
	}

	return res
}
