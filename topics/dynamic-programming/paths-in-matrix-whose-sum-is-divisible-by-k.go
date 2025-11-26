package dynamicprogramming

// link: https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
// thought: 数量问题，初始化和取模变成 m×n×k 的时间复杂度的问题
func numberOfPaths(grid [][]int, k int) int {
	const MOD = 1e9 + 7
	m, n := len(grid), len(grid[0])

	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				dp[i][j][grid[0][0]%k] = 1
				continue
			}

			v := grid[i-1][j-1] % k
			for r := 0; r < k; r++ {
				prev := (r - v + k) % k
				dp[i][j][r] = (dp[i-1][j][prev] + dp[i][j-1][prev]) % MOD
			}
		}
	}

	return dp[m][n][0]
}
