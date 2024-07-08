package Array

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, points[0])

	for i := 1; i < len(points); i++ {
		current := points[i]
		last := merged[len(merged)-1]

		if current[0] > last[1] {
			merged = append(merged, current)
		} else {
			// 更新 last 区间的结束位置为 last 和 current 的结束位置的较大值
			last[1] = max(last[1], current[1])
			merged[len(merged)-1] = last
		}
	}

	return len(merged)
}
