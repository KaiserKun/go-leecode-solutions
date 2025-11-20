package ps

// link: https://leetcode.cn/problems/find-pivot-index/
// thought: 前缀和，左侧元素和就是前缀和，右侧元素和为prefix[n]-prefix[i+1]
func pivotIndex(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	prefix[0] = 0
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	for i := 0; i < n; i++ {
		if prefix[i] == prefix[n]-prefix[i+1] {
			return i
		}
	}
	return -1
}
