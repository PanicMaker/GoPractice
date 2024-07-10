package HashMap

// https://leetcode.cn/problems/happy-number/description/

// 通过一个哈希表记录数字是否重复出现
func isHappyI(n int) bool {
	step := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return sum
	}

	maps := make(map[int]bool)

	for n != 1 && !maps[n] {
		n, maps[n] = step(n), true
	}

	return n == 1
}

// 通过快慢指针判断，若不是快乐数，则会形成一个环
func isHappyII(n int) bool {
	step := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return sum
	}

	slow, fast := n, step(n)

	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}

	return fast == 1
}
