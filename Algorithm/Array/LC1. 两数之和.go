package Array

// https://leetcode.cn/problems/two-sum/description/

// 两个循环
func twoSumI(nums []int, target int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 使用哈希表
func twoSumII(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))

	for i, v := range nums {
		if val, ok := maps[target-v]; ok {
			return []int{i, val}
		}
		maps[v] = i
	}

	return nil
}
