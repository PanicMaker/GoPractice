package Array

import "slices"

// 以k为分界线拼接，最后复制到原切片中
func rotateI(nums []int, k int) {
	k %= len(nums)
	length := len(nums)

	copy(nums, append(nums[length-k:], nums[:length-k]...))
}

// 翻转
func rotateII(nums []int, k int) {
	k %= len(nums)

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

// 用额外数组
func rotateIII(nums []int, k int) {
	newNums := make([]int, len(nums))

	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}
