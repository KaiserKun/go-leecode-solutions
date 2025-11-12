package medium

// link: https://leetcode.cn/problems/open-the-lock/
// thought: 双向BFS扩展
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	// map记录deadends
	dead := make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	}
	if dead["0000"] {
		return -1
	}

	beginQueue := []string{"0000"}
	endQueue := []string{target}
	beginSeen := make(map[string]int)
	endSeen := make(map[string]int)
	beginSeen["0000"] = 0
	endSeen[target] = 0

	dir := []int{1, -1}

	for len(beginQueue) > 0 && len(endQueue) > 0 {
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
			beginSeen, endSeen = endSeen, beginSeen
		}

		size := len(beginQueue)
		for i := 0; i < size; i++ {
			cur := beginQueue[0]
			beginQueue = beginQueue[1:]
			for j := 0; j < 4; j++ {
				for _, v := range dir {
					var newCh byte
					if v == 1 {
						// 向上旋转
						if cur[j] == '9' {
							newCh = '0'
						} else {
							newCh = cur[j] + 1
						}
					} else {
						// 向下旋转
						if cur[j] == '0' {
							newCh = '9'
						} else {
							newCh = cur[j] - 1
						}
					}
					newS := cur[:j] + string(newCh) + cur[j+1:]
					if dead[newS] {
						continue
					}
					if steps, ok := endSeen[newS]; ok {
						return beginSeen[cur] + steps + 1
					}

					if beginSeen[newS] == 0 {
						beginQueue = append(beginQueue, newS)
						beginSeen[newS] = beginSeen[cur] + 1
					}
				}
			}
		}
	}
	return -1
}
