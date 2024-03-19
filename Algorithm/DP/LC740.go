package DP

func deleteAndEarn1(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}

	arr := make([]int, maxNum+1)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += nums[i]
	}

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[len(arr)-1]
}

func deleteAndEarn2(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}

	arr := make([]int, maxNum+1)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += nums[i]
	}

	pre, cur := arr[0], max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		pre, cur = cur, max(cur, pre+arr[i])

	}
	return cur
}
