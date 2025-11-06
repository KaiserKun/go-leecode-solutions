package medium

// link: https://leetcode.cn/problems/01-matrix/
// thought: 多源BFS，从0的位置出发
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	ans := make([][]int, m)
	q := [][]int{}
	for i := range ans {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
				ans[i][j] = 0
			} else {
				ans[i][j] = 1<<31 - 1
			}
		}
	}
	direct := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]

		for _, dir := range direct {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if ans[x][y] > ans[i][j]+1 {
					ans[x][y] = ans[i][j] + 1
					q = append(q, []int{x, y})
				}
			}
		}
	}
	return ans
}
