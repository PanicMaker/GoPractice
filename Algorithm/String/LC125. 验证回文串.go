package String

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// https://leetcode.cn/problems/valid-palindrome/

func isPalindromeI(s string) bool {
	s = strings.ToLower(s)
	pattern := regexp.MustCompile(`[^a-z0-9]+`)
	s = pattern.ReplaceAllString(s, "")

	fmt.Println(s)

	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	left, right := 0, n-1

	for left < right {
		// 跳过非字母和数字的字符
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}
