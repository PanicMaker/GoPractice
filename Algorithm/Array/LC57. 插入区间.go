package Array

import (
	"sort"
)

// https://leetcode.cn/problems/insert-interval

// 将 newInterval 加入 intervals 后，用 LC56 的合并算法做
func insertI(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := merged[len(merged)-1]

		if current[0] > last[1] {
			merged = append(merged, current)
		} else {
			// 更新 last 区间的结束位置为 last 和 current 的结束位置的较大值
			last[1] = max(last[1], current[1])
			merged[len(merged)-1] = last
		}
	}

	return merged
}

func insertII(intervals [][]int, newInterval []int) [][]int {
	merged := make([][]int, 0)
	i := 0
	n := len(intervals)

	// 将没有重叠的区间先加入结果
	for i < n && intervals[i][1] < newInterval[0] {
		merged = append(merged, intervals[i])
		i++
	}

	// 处理有重叠的区间
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	merged = append(merged, newInterval)

	// 将剩下的区间加入结果
	for i < n {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}
