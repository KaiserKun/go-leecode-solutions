package easy

// link: https://leetcode.cn/problems/find-the-k-th-character-in-string-game-i/
// thought:
import "math/bits"

func kthCharacter(k int64) byte {
	ans := 0
	t := 0
	for k != 1 {
		t = bits.Len64(uint64(k)) - 1
		if (1 << t) == k {
			t--
		}
		k = k - (1 << t)
		ans++

	}
	return byte('a' + (ans % 26))
}
