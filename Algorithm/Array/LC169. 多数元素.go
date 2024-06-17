package Array

import "sort"

// https://leetcode.cn/problems/majority-element/description/

// 使用一个map统计数字出现次数，时间复杂度O(N)，空间复杂度O(N)
func majorityElementI(nums []int) int {
	dict := make(map[int]int, len(nums))
	maxCount := 0
	result := nums[0]

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			dict[num]++
		} else {
			dict[num]++
		}
	}

	for k, v := range dict {
		maxCount = max(maxCount, v)
		if v == maxCount {
			result = k
		}
	}

	return result
}

// 索引为 len(nums)/2 的数一定是符合条件的
func majorityElementII(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore 投票算法
// 把众数记为 +1+1+1，把其他数记为 −1-1−1，将它们全部加起来，显然和大于 0
func majorityElementIII(nums []int) int {
	val := nums[0]
	count := 0

	for _, num := range nums {
		if count == 0 {
			val = num
		}

		if num == val {
			count++
		} else {
			count--
		}
	}

	return val
}
