package Array

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
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
