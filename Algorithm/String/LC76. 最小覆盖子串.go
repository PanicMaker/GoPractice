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
	need := make(map[byte]int)   // 用于统计需要凑齐的字符
	window := make(map[byte]int) // 记录滑动窗口内已有字符的个数

	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0               // 滑动窗口的左右边界
	valid := 0                        // 判断窗口中是否已经包含了字串 t 中所有字符
	start, length := 0, math.MaxInt32 // 最小覆盖子串的起始索引及长度

	for right < len(s) { // 当 right 小于 s 的长度时，继续循环
		c := s[right] // c 是将要加入窗口中的字符
		right++
		if _, ok := need[c]; ok { // 如果这个字符在字串 t 中需要的话
			window[c]++               // 加入窗口中
			if window[c] == need[c] { // 如果字符 c 在窗口中的数量已经满足其在字串 t 中的数量
				valid++ // 计数器 valid 加一
			}
		}
		for valid == len(need) { // 如果滑动窗口中的字符已经完全覆盖字串 t 中的字符
			if right-left < length { // 如果此时的覆盖子串更短
				start = left          // 更新最小覆盖子串的起始索引
				length = right - left // 更新最小子串的长度
			}
			d := s[left]              // d 是将要移出窗口的字符
			left++                    // 左侧窗口右移
			if _, ok := need[d]; ok { // 如果这个字符在字串 t 中需要的话
				if window[d] == need[d] { // 如果这个字符已经满足了他在字串 t 中的需求
					valid-- // 计数器 valid 减一
				}
				window[d]-- // 移出窗口
			}
		}
	}
	if length == math.MaxInt32 { // 如果最小子串长度没有更新，则返回空格
		return ""
	}
	return s[start : start+length] // 返回最小覆盖子串
}
