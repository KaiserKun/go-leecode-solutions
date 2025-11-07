package bfs

// link: https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/
// thought: 使用淹没法优化，本质上是一个多源BFS
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])

	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	q := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = '+'
	ans := 0

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			curx, cury := q[0][0], q[0][1]
			q = q[1:]
			for _, dir := range direction {
				x, y := curx+dir[0], cury+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == '.' {
					if x == 0 || x == m-1 || y == 0 || y == n-1 {
						return ans + 1
					}
					maze[x][y] = '+'
					q = append(q, []int{x, y})
				}
			}
		}
		ans++
	}
	return -1

}
