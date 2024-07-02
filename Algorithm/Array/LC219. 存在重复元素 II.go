package Array

// https://leetcode.cn/problems/contains-duplicate-ii/

func containsNearbyDuplicate(nums []int, k int) bool {
	// 使用 map 记录元素及其最后出现的位置
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// 检查当前元素是否已经在 map 中出现过
		if idx, ok := hash[nums[i]]; ok {
			// 如果已经出现过，检查位置差是否 <= k
			if i-idx <= k {
				return true
			}
		}
		// 更新元素的最后出现位置为当前位置 i
		hash[nums[i]] = i
	}

	return false
}
