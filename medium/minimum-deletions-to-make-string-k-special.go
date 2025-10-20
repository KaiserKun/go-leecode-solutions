package medium

// link: https://leetcode.cn/problems/minimum-deletions-to-make-string-k-special/
// thought: 满足 k-special 的条件是每个字符的出现次数差不超过 k，所以当以a作为下限时，所有字母出现的次数都在[a, a+k]范围内
func minimumDeletions(word string, k int) int {
	var cnt [26]int
	// 记录每个字符的出现次数
	for _, c := range word {
		cnt[c-'a']++
	}
	res := len(word)
	// 用a作为下限，计算删除的最小次数
	for _, a := range cnt {
		del := 0
		for _, b := range cnt {
			// 当出现的字符串比a小时，不满足k-special，需全部删除
			if a > b {
				del += b
			} else if b > a+k {
				del += b - a - k
			}
		}
		res = min(res, del)
	}
	return res
}
