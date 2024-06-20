package Array

func rotateI(nums []int, k int) {
	k %= len(nums)
	length := len(nums)

	copy(nums, append(nums[length-k:], nums[:length-k]...))

	return
}
