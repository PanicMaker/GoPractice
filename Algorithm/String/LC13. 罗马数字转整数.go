package String

// https://leetcode.cn/problems/roman-to-integer/description

// 当前位置的元素比下个元素小，就减去当前值，否则加上当前值
func romanToIntI(s string) int {
	datas := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && datas[s[i]] < datas[s[i+1]] {
			sum -= datas[s[i]]
		} else {
			sum += datas[s[i]]
		}
	}

	return sum
}

func romanToIntII(s string) int {
	datas := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	high := 1
	// 记录当前遇到的最大数字，更大就加，小则减
	for i := len(s) - 1; i >= 0; i-- {
		num := datas[s[i]]
		if num >= high {
			sum += num
			high = num
		} else {
			sum -= num
		}
	}

	return sum
}
