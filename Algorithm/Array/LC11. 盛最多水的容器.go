package Array

// https://leetcode.cn/problems/container-with-most-water

// 双循环，时间复杂度O(n^2)，题目超时
func maxAreaI(height []int) int {
	if len(height) == 2 {
		return min(height[0], height[1])
	}

	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := min(height[i], height[j])
			multiply := h * (j - i)
			res = max(res, multiply)
		}
	}
	return res
}

// 双指针解法
func maxAreaII(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
