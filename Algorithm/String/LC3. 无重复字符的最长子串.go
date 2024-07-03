package String

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	res := 0

	left, right := 0, 0

	for right < len(s) {
		c := s[right]
		right++

		window[c]++

		// 大于1时窗口内有重复字符
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		res = max(res, right-left)
	}

	return res
}
