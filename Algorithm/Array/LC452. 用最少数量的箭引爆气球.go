package Array

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按照每个气球的右边界进行排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	index := points[0][1]

	for i := 0; i < len(points); i++ {
		// 如果当前气球的左边界大于当前箭的射击位置，需要增加一支箭
		if points[i][0] > index {
			count++
			index = points[i][1]
		}

	}

	return count
}
