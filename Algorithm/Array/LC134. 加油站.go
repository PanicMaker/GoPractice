package Array

// https://leetcode.cn/problems/gas-station/

// 时间复杂度为 O(n)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	totalGas := 0
	totalCost := 0
	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1 // 总的油量不足以完成一圈
	}

	start := 0
	remain := 0

	for i := 0; i < n; i++ {
		remain += gas[i] - cost[i]
		if remain < 0 {
			start = i + 1 // 更新起始点为下一个点
			remain = 0    // 重置剩余油量
		}
	}

	return start
}
