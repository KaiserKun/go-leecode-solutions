package hard

// link: https://leetcode.cn/problems/maximum-candies-you-can-get-from-boxes/
// thought: BFS的一个使用情况
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	ans := 0
	has := make([]bool, n)
	opened := make([]bool, n)
	queue := []int{}

	for _, v := range initialBoxes {
		has[v] = true
		if status[v] == 1 {
			queue = append(queue, v)
		}
	}
	for len(queue) > 0 {
		box := queue[0]
		queue = queue[1:]

		if opened[box] {
			continue
		}
		opened[box] = true

		ans += candies[box]

		// 拿到钥匙
		for _, v := range keys[box] {
			if status[v] == 0 {
				// 标记可以打开
				status[v] = 1
				// 有箱子而且没有打开，再加入队列
				if has[v] && !opened[v] {
					queue = append(queue, v)
				}
			}
		}

		// 拿到其中的箱子
		for _, v := range containedBoxes[box] {
			if !has[v] {
				has[v] = true
				// 可以打开，但是没有打开
				if status[v] == 1 && !opened[v] {
					queue = append(queue, v)
				}
			}
		}
	}
	return ans
}
