package easy

// link: https://leetcode.cn/problems/power-of-four/
// thought: 递归/循环，不断除以4，直到结果为1或无法整除

func isPowerOfFour(n int) bool {
	for n > 0 && n%4 == 0 {
		n /= 4
	}
	return n == 1
}
