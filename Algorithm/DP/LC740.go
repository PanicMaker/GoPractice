package DP

import "fmt"

func deleteAndEarn(nums []int) int {
	arr := make([]int, 10000)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += nums[1]
	}
	fmt.Println(arr)

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[len(arr)-1]
}
