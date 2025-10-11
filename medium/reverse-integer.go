package medium

// link: https://leetcode.com/problems/reverse-integer/
// thought: 利用数学方法反转整数
func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}
	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	// 负数反转之后，只有可能比 1<<31-1 大，并不会比 -(1<<31) 小
	if ret > 1<<31-1 {
		return 0
	}

	return ret * flag
}
