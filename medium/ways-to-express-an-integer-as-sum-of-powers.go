package medium

// link: https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/
// thought: 0/1 背包计数问题，dp[j]表示和为j的方案数，外层循环物品，内层循环背包容量

const mod = 1000000007

func powInt(base, exp, limit int) (int, bool) {
	// 返回 base^exp，第二个返回值表示是否超过 limit（或溢出）
	res := 1
	for k := 0; k < exp; k++ {
		res *= base
		if res > limit {
			return 0, true
		}
	}
	return res, false
}

func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	// 枚举物品（幂次）
	for i := 1; ; i++ {
		// 计算 i^x，如果超过 n 则停止
		val, over := powInt(i, x, n)
		if over || val > n {
			break
		}
		// 倒序 -> 0/1 背包
		for j := n; j >= val; j-- {
			dp[j] = (dp[j] + dp[j-val]) % mod
		}
	}
	return dp[n]
}
