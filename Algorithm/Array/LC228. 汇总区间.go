package Array

import (
	"fmt"
)

// https://leetcode.cn/problems/summary-ranges/

func summaryRanges(nums []int) []string {
	n := len(nums)
	var result []string

	if n == 0 {
		return result
	}

	i := 0

	for j := 1; j <= n; j++ {
		if j < n && nums[j] == nums[j-1]+1 {
			continue
		}
		if j-i > 1 {
			str := fmt.Sprintf("%d->%d", nums[i], nums[j-1])
			result = append(result, str)
		} else {
			str := fmt.Sprintf("%d", nums[i])
			result = append(result, str)
		}
		i = j
	}

	return result
}
