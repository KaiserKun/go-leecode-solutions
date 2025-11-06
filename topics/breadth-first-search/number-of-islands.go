package bfs

// link: https://leetcode.cn/problems/number-of-islands/
// thought: BFS遍历遇到的岛屿，使用一个切片记录现在的值是否已经访问过
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	ans := 0
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				q := [][]int{{i, j}}
				visited[i][j] = true

				for len(q) > 0 {
					curX, curY := q[0][0], q[0][1]
					q = q[1:]
					for _, dir := range direct {
						x, y := curX+dir[0], curY+dir[1]
						if x >= 0 && y >= 0 && x < m && y < n &&
							grid[x][y] == '1' && !visited[x][y] {
							q = append(q, []int{x, y})
							visited[x][y] = true
						}
					}
				}
			}
		}
	}
	return ans
}
