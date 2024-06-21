package Array

// https://leetcode.cn/problems/product-of-array-except-self/description/

// 两个循环，时间复杂度为O(n²)，题目超时
func productExceptSelfI(nums []int) []int {
	plus := func(nums []int) int {
		result := 1

		for _, v := range nums {
			result *= v
		}
		return result
	}

	var result []int
	for i, _ := range nums {
		if i == 0 {
			result = append(result, plus(nums[i+1:]))
		} else if i == len(nums)-1 {
			result = append(result, plus(nums[:len(nums)-1]))
		} else {
			result = append(result, plus(nums[:i])*plus(nums[i+1:]))
		}
	}
	return result
}

// 时间复杂度为𝑂(𝑛)，空间复杂度为𝑂(1)
func productExceptSelfII(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	// 从左到右计算左乘积
	left := 1
	for i := 0; i < n; i++ {
		result[i] *= left
		left *= nums[i]
		// fmt.Printf("result = %v, left = %d \n", result, left)
	}

	// 从右到左计算右乘积，并将左乘积与右乘积相乘
	right := 1
	for j := n - 1; j >= 0; j-- {
		result[j] *= right
		right *= nums[j]
		// fmt.Printf("result = %v, right = %d \n", result, right)
	}

	return result
}
