package medium

// link: https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/
// thought: 贪心算法，使用一个最小堆来存储当前可以参加的活动的结束时间
import "sort"

func maxEvents(events [][]int) int {
	n := len(events)
	// 计算最大的结束时间
	maxDay := 0
	for i := 0; i < len(events); i++ {
		maxDay = max(maxDay, events[i][1])
	}
	// 将活动按开始时间排序
	pq := make([]int, 0)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	ans := 0
	for i, j := 0, 0; i <= maxDay; i++ {
		for j < n && events[j][0] <= i {
			pq = append(pq, events[j][1])
			j++
		}
		for len(pq) > 0 && pq[0] < i {
			pq = pq[1:]
		}
		if len(pq) > 0 {
			pq = pq[1:]
			ans++
		}
	}

	return ans
}
