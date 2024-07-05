package String

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string/description/

// 双指针
func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
