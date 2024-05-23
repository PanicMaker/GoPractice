package main

// 852. 山脉数组的峰顶索引
// https://leetcode.cn/problems/peak-index-in-a-mountain-array/description/

// 双指针遍历左右子数组，容易超时
func peakIndexInMountainArray1(arr []int) int {
	for i := 1; i < len(arr); i++ {
		l1, r1 := i, i
		l2 := l1 - 1
		for l2 >= 0 {
			if arr[l2] < arr[l1] {
				l2--
				l1--
			} else {
				break
			}
		}

		r2 := r1 + 1
		for r2 < len(arr) {
			if arr[r2] < arr[r1] {
				r2++
				r1++
			} else {
				break
			}
		}

		if l1 == 0 && r1 == len(arr)-1 {
			return i
		}
	}
	return -1
}

// 使用二分查找解答
func peakIndexInMountainArray2(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := (left + right) / 2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else if arr[mid] > arr[mid+1] {
			right = mid
		}
	}
	return left
}
