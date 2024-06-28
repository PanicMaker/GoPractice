package String

import "math"

// https://leetcode.cn/problems/minimum-window-substring/

// 双重循环查找每一个子串是否包含目标字符，会超时
func minWindowI(s string, t string) string {
	var result string
	n := len(s)
	minLength := n

	containsAll := func(str, t string) bool {
		count := make(map[rune]int)
		for _, ch := range t {
			count[ch]++
		}

		for _, ch := range str {
			if count[ch] > 0 {
				count[ch]--
			}
		}

		for _, v := range count {
			if v > 0 {
				return false
			}
		}

		return true
	}

	for i := 0; i < n; i++ {
		for j := i; j <= n; j++ {
			str := s[i:j]
			if containsAll(str, t) {
				if len(str) <= minLength {
					minLength = len(str)
					result = str
				}
			}
		}
	}

	return result
}

// 用滑动窗口算法
func minWindowII(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start, length := 0, math.MaxInt

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}

	return s[start : start+length]
}
