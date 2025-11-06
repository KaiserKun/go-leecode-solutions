package medium

// link: https://leetcode.cn/problems/rotting-oranges/
// thought: 多源bfs，利用腐烂的橘子作为起点，同样是水波状扩散，层级扩展
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				cnt++
			}
		}
	}

	if cnt == 0 {
		return 0
	}

	direct := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := 0

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curx, cury := q[0][0], q[0][1]
			q = q[1:]

			for _, dir := range direct {
				x, y := curx+dir[0], cury+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					cnt--
					q = append(q, []int{x, y})
				}
			}
		}
		if len(q) > 0 {
			ans++
		}
	}

	if cnt > 0 {
		return -1
	}

	return ans
}
