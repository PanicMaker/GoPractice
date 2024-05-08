package LinkedList

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return len(nums[:i+1])
}
