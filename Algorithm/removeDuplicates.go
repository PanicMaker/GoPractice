package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	left := 0
	for right := 1; right < len(nums); right++ {
		//如果左指针和右指针指向的值一样，说明有重复的，
		//这个时候，左指针不动，右指针继续往右移。如果他俩
		//指向的值不一样就把右指针指向的值往前挪
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

func main() {
	nums := [...]int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3}
	fmt.Print(removeDuplicates(nums[:]))
}
