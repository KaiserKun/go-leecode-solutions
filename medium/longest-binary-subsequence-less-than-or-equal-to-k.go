package medium

import "math/bits"

// link: https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// thought:

func longestSubsequence(s string, k int) int {
	var sm int64 = 0
	cnt := 0

	// bitLen is the number of bits needed to represent k (floor(log2(k))+1)
	bitLen := bits.Len(uint(k))
	k64 := int64(k)

	for i := 0; i < len(s); i++ {
		ch := s[len(s)-1-i]
		if ch == '1' {
			if i < bitLen {
				val := int64(1) << uint(i)
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
