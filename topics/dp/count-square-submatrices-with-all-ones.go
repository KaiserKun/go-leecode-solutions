package dp

// link: https://leetcode.cn/problems/count-square-submatrices-with-all-ones/
// thought: 动态规划，dp[i][j]表示以(i,j)为右下角的全1正方形的边长
func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化第一行和第一列
	dp[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = matrix[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				// 当前位置为1，取左、上、左上三个位置的最小值加1
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
			// 累加所有位置的dp值
			ans += dp[i][j]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
