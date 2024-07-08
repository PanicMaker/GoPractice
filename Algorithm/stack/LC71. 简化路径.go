package stack

import "strings"

// https://leetcode.cn/problems/simplify-path

func simplifyPath(path string) string {
	stack := make([]string, 0)
	splits := strings.Split(path, "/")

	for _, v := range splits {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}
	}

	res := "/" + strings.Join(stack, "/")

	return res
}
