package Array

import "sort"

// https://leetcode.cn/problems/3sum/description

func threeSumI(nums []int) [][]int {
	n := len(nums)
	var result [][]int

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1
		target := 0 - nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func threeSumII(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	uniqueTriplets := make(map[[3]int]struct{})

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// 跳过重复的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		maps := make(map[int]int, len(nums))
		for x := i + 1; x < n; x++ {
			if val, ok := maps[target-nums[x]]; ok {
				triplet := [3]int{nums[i], nums[x], nums[val]}
				// 使用map的key不可重复来判断答案是否重复
				if _, found := uniqueTriplets[triplet]; !found {
					result = append(result, []int{nums[i], nums[x], nums[val]})
					uniqueTriplets[triplet] = struct{}{}
				}
			}
			maps[nums[x]] = x
		}
	}

	return result
}
