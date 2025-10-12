package dp

// link: https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/
// thought: 动态规划，dp[i][j]表示前i个活动中选择j个活动的最大价值
import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(events [][]int, k int) int {
	n := len(events)
	dp := make([][]int, n+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	// 按结束时间排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	for i := 0; i < n; i++ {
		// 二分查找前一个不冲突的活动
		p := sort.Search(i, func(j int) bool {
			return events[j][1] >= events[i][0]
		})
		for j := 1; j <= k; j++ {
			dp[i+1][j] = max(dp[i][j], dp[p][j-1]+events[i][2])
		}
	}

	return dp[n][k]
}
