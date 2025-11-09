package bfs

// link: https://leetcode.cn/problems/sliding-puzzle/
// thought: 分层 + 状态判断
func slidingPuzzle(board [][]int) int {
	if judge(board) {
		return 0
	}

	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	ans := 0
	q := [][][]int{board}
	visited := make(map[string]bool)
	visited[getKey(board)] = true

	for len(q) > 0 {
		size := len(q)
		ans++

		for i := 0; i < size; i++ {
			board = q[0]
			q = q[1:]

			// 找到0板块的位置
			curx, cury := -1, -1
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					if board[i][j] == 0 {
						curx, cury = i, j
					}
				}
			}
			for _, dir := range direction {
				x, y := curx+dir[0], cury+dir[1]
				if x < 0 || x >= 2 || y < 0 || y >= 3 {
					continue
				}

				tmp := make([][]int, 2)
				for i := 0; i < 2; i++ {
					tmp[i] = make([]int, 3)
					copy(tmp[i], board[i])
				}
				tmp[x][y], tmp[curx][cury] = tmp[curx][cury], tmp[x][y]

				key := getKey(tmp)
				if !visited[key] {
					if judge(tmp) {
						return ans
					}
					visited[key] = true
					q = append(q, tmp)
				}
			}
		}

	}
	if len(q) == 0 {
		return -1
	}
	return ans
}

func getKey(board [][]int) string {
	key := ""
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			key += string('0' + board[i][j])
		}
	}
	return key
}

func judge(board [][]int) bool {
	tmp := [][]int{{1, 2, 3}, {4, 5, 0}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != tmp[i][j] {
				return false
			}
		}
	}
	return true
}
