package Array

import (
	"math"
)

// https://leetcode.cn/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	res := math.MaxInt

	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			res = min(res, right-left)
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}
