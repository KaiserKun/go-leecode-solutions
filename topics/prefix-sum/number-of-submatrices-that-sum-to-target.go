package ps

// link: https://leetcode.cn/problems/number-of-submatrices-that-sum-to-target/
// thought: 要确定是矩阵，且sumcnt哈希计算的值可用的情况下，只能在同一行，不用行得出的差是一个异形形状
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])

	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]
		}
	}

	ans := 0
	for c1 := 0; c1 < n; c1++ {
		for c2 := c1; c2 < n; c2++ {
			// 列范围固定后，j 就固定了
			// 现在只在行方向变化，所以差一定是矩形
			sumcnt := make(map[int]int)
			sumcnt[0] = 1

			for i := 0; i < m; i++ {
				sum := prefix[i+1][c2+1] - prefix[i+1][c1]
				if cnt, ok := sumcnt[sum-target]; ok {
					ans += cnt
				}
				sumcnt[sum]++
			}
		}
	}

	return ans
}
