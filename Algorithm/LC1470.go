package main

// 1470. 重新排列数组
// https://leetcode.cn/problems/shuffle-the-array/description/

func shuffle1(nums []int, n int) []int {
	nums1, nums2 := nums[:n], nums[n:]
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, nums1[i], nums2[i])
	}
	return res
}

func shuffle2(nums []int, n int) []int {
	var res []int
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		res = append(res, nums[i], nums[j])
	}
	return res
}
