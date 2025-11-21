package medium

// link: https://leetcode.cn/problems/range-sum-query-2d-immutable/
// thought: 标准前缀和统计每行的前i个元素和，不同行叠加得到矩形和

type NumMatrix struct {
	matrix [][]int
	prefix [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	prefix := make([][]int, m)
	for i := range prefix {
		prefix[i] = make([]int, n)
	}
	// 先加行
	for i := 0; i < m; i++ {
		prefix[i][0] = matrix[i][0]

	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			prefix[i][j] = prefix[i][j-1] + matrix[i][j]
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i][j] += prefix[i-1][j]
		}
	}

	return NumMatrix{matrix, prefix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	total := this.prefix[row2][col2]

	if row1 > 0 {
		total -= this.prefix[row1-1][col2]
	}
	if col1 > 0 {
		total -= this.prefix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		total += this.prefix[row1-1][col1-1]
	}

	return total
}
