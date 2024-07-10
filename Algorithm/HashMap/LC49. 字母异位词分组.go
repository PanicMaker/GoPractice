package HashMap

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/group-anagrams

// 双重循环，跳过已分类的元素，依然会超时
func groupAnagramsI(strs []string) [][]string {
	// 初始化结果数组
	res := make([][]string, 0)

	// 如果输入数组只有一个字符串，直接将其作为一个分组返回
	if len(strs) == 1 {
		res = append(res, strs)
		return res
	}

	// 标记数组，用于记录每个字符串是否已经被归类到某个分组
	flag := make(map[int]int)
	for i := range strs {
		flag[i] = 0 // 初始化标记，0 表示未归类，1 表示已归类
	}

	// 遍历每个字符串进行分组操作
	for i := 0; i < len(strs); i++ {
		si := strs[i] // 当前字符串 si

		// 如果当前字符串已经归类过，则跳过
		if flag[i] == 1 {
			continue
		}

		// 临时数组，存储当前字母异位词分组的所有字符串
		tmp := make([]string, 0)
		tmp = append(tmp, si) // 将当前字符串 si 添加到临时数组中
		flag[i]++             // 标记当前字符串 si 已归类

		// 遍历后续的字符串，查找与当前字符串 si 是字母异位词的字符串
		for j := i + 1; j < len(strs); j++ {
			// 如果字符串已经归类过，则跳过
			if flag[j] == 1 {
				continue
			}

			sj := strs[j] // 当前字符串 sj

			// 判断 si 和 sj 是否是字母异位词
			is := true

			// 如果长度不相等，直接跳过
			if len(si) != len(sj) {
				continue
			}

			// 使用 map 计数判断是否为字母异位词
			cnt := make(map[rune]int)
			for _, v := range si {
				cnt[v]++
			}

			for _, v := range sj {
				cnt[v]--
				if cnt[v] < 0 {
					is = false
					break
				}
			}

			// 如果是字母异位词，将 sj 添加到当前分组的临时数组中，并标记已归类
			if is {
				tmp = append(tmp, sj)
				flag[j]++
			}
		}

		// 将当前分组的临时数组添加到结果数组中
		res = append(res, tmp)
	}

	return res
}

func groupAnagramsII(strs []string) [][]string {
	// 声明一个 map 用于存储分组后的结果
	groups := make(map[string][]string)

	// 遍历每个字符串
	for _, str := range strs {
		// 将字符串转换为字符数组，便于排序
		chars := strings.Split(str, "")
		// 对字符数组进行排序
		sort.Strings(chars)
		// 排序后的字符串作为哈希表的键
		sortedStr := strings.Join(chars, "")

		// 将原始字符串添加到对应的分组中
		groups[sortedStr] = append(groups[sortedStr], str)
	}

	// 将 map 中的值转换为二维数组返回
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
