package medium

// link: https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// thought: 对nums中的每个元素进行取模运算，压缩到[0,value)范围内，记录每个余数出现的次数，随后逐次递增，直到找到未出现的数
func findSmallestInteger(nums []int, value int) int {
	// 记录出现过的数
	cnt := make([]int, value)
	for _, num := range nums {
		// 取模运算，压缩到[0,value)范围内
		if num >= 0 {
			cnt[num%value]++
		}
	}
	// 找到最小的未出现的非负整数
	m := 0
	for {
		r := m % value
		if cnt[r] == 0 {
			return m
		}
		cnt[r]--
		m++
	}
}
