package easy

// link: https://leetcode.com/problems/find-lucky-integer-in-an-array/
// thought: 使用哈希表统计元素出现频率
func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	ans := -1
	for k, v := range m {
		if k == v {
			ans = max(ans, k)
		}
	}
	return ans
}
