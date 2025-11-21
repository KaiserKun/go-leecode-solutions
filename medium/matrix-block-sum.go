package medium

// link: https://leetcode.cn/problems/matrix-block-sum/
// thought: 本质是，以下标i,j为中心，长为2k+1的正方形
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i+1][j] + mat[i][j]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefix[i][j] += prefix[i-1][j]
		}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r1, c1 := i-k, j-k
			r2, c2 := i+k, j+k
			if i+k >= m {
				r2 = m - 1
			}
			if j+k >= n {
				c2 = n - 1
			}
			if i-k < 0 {
				r1 = 0
			}
			if j-k < 0 {
				c1 = 0
			}
			ans[i][j] = prefix[r2+1][c2+1] + prefix[r1][c1] - prefix[r2+1][c1] - prefix[r1][c2+1]
		}
	}

	return ans
}
