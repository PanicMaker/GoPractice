package Array

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/

// 双指针该类型题的通用解法
func removeDuplicates80I(nums []int) int {
	//每个元素最多出现的次数
	k := 2

	if len(nums) <= k {
		return len(nums)
	}

	// 初始化快慢指针
	i, j := k, k

	for j < len(nums) {
		// 如果 nums[j] 不等于 nums[i-k]
		// 则将 nums[j] 复制到 nums[i], i++
		if nums[i-k] != nums[j] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}

// 双指针增加一个中间变量统计数出现的次数
func removeDuplicates80II(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	i := 0
	count := 1

	for j := 1; j < len(nums); j++ {
		// 当前数字和前一个数字相同，增加计数器
		if nums[j] == nums[j-1] {
			count++
		} else {
			count = 1
		}

		// 检查当前数字的出现次数是否小于等于2（if count <= 2）
		// 如果count小于等于2，说明当前数字的出现次数没有超过2次，可以将其添加到新数组中
		// 使用指针j指向新数组的位置，将当前数字（nums[j]）复制到新数组的位置（nums[i]）
		// 将指针i右移一位，为下一个符合条件的数字腾出位置
		if count <= 2 {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

// 用一个map记录了数组中各个数出现的次数并在一个新数组中添加，不符合题目要求
func removeDuplicates80III(nums []int) int {
	maps := make(map[int]int, len(nums))
	var newNums []int

	for _, num := range nums {
		v, ok := maps[num]
		if !ok || v < 2 {
			newNums = append(newNums, num)
			maps[num]++
		}
	}

	copy(nums, newNums)
	return len(newNums)
}
