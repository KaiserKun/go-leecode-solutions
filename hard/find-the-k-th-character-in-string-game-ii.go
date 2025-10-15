package hard

import "math/bits"

// link: https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/
// thought: 通过位运算和操作次数来确定第k个字符
func kthCharacter(k int64, operations []int) byte {
	ans := 0
	t := 0
	for k != 1 {
		t = bits.Len64(uint64(k)) - 1
		if (1 << t) == k {
			t--
		}
		k = k - (1 << t)
		if operations[t] > 0 {
			ans++
		}
	}
	return byte('a' + (ans % 26))
}
