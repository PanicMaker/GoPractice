package appendInt

func appendInt(arr []int, v int) []int {
	z := []int{}
	zlen := len(arr) + 1
	if zlen <= cap(arr) {
		z = z[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(arr) {
			zcap = 2 * len(arr)
		}
		z = make([]int, zlen, zcap)
		copy(z, arr)
	}
	z[len(arr)] = v
	return z
}
