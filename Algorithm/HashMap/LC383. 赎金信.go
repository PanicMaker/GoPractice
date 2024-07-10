package HashMap

// https://leetcode.cn/problems/ransom-note

func canConstruct(ransomNote string, magazine string) bool {
	hash := make(map[rune]int)

	for _, v := range magazine {
		hash[v]++
	}

	for _, v := range ransomNote {
		hash[v]--
		if hash[v] < 0 {
			return false
		}
	}

	return true
}
