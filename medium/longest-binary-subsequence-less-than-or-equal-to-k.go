package medium

import "math/bits"

// link: https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// thought: 贪心，从低位向高位选择'0'，再选择能加上不超过k的'1'
func longestSubsequence(s string, k int) int {
	var sm int64 = 0
	cnt := 0

	// bitlen是k的二进制长度
	bitLen := bits.Len(uint(k))
	// k64是k的64位整数表示
	k64 := int64(k)
	// 从低位向高位遍历字符串
	for i := 0; i < len(s); i++ {
		// 从字符串末尾开始取字符
		ch := s[len(s)-1-i]
		// 处理'1'字符
		if ch == '1' {
			// 只考虑k的二进制长度范围内的'1'
			if i < bitLen {
				// 计算当前位的值
				val := int64(1) << uint(i)
				// 检查加上当前位后是否不超过k
				if sm+val <= k64 {
					sm += val
					cnt++
				}
			}
		} else {
			cnt++
		}
	}
	return cnt
}
