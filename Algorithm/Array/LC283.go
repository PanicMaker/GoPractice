package Array

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/description/

func moveZeroes(nums []int) {
	slow, fast := 0, 0

	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}
