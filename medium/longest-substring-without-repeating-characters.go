package medium

// link: https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// thought: 使用滑动窗口和哈希表记录字符的最新索引，动态调整窗口的起始位置以确保子串无重复字符
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	// 遍历字符串，更新窗口起始位置和最大长度
	for i := 0; i < len(s); i++ {
		if lastIndex, ok := charIndex[s[i]]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndex[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
