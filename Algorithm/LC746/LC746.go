package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// 将顶部当做没有任何花费的台阶计入
	cost = append(cost, 0)

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func main() {
	cost1 := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost1))
	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost2))
}
