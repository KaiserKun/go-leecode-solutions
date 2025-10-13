package hard

// link: https://leetcode.cn/problems/regular-expression-matching/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[i][j]表示s的前i个字符和p的前j个字符是否匹配
	dp := make([][]bool, m+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true

	// 从后向前填表
	for i := m; i >= 0; i-- {
		// j从后向前遍历模式串
		for j := n - 1; j >= 0; j-- {
			firstMatch := i < m && (s[i] == p[j] || p[j] == '.')
			// 处理 '*' 情况，
			if j+1 < n && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
