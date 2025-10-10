package medium

// link: https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three/
// thought: 循环除以3，检查每一位是否为2，因为n若可以表示为3的幂的和，对3取模之后只能是0或1
func checkPowersOfThree(n int) bool {
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}
