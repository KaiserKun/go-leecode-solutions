package easy

// link: https://leetcode.cn/problems/power-of-three/
// thought: 若n是3的幂，则不断除以3后最终结果为1，且过程中不能出现余数
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
