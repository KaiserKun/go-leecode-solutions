package hard

// link: https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/
// thought: 按前缀树组织，重点在于计算当前子树的数量
func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := calc(n, cur, cur+1)
		if steps <= k {
			cur++
			k -= steps
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func calc(n, cur, next int) int {
	steps := 0
	for cur <= n {
		steps += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}
	return steps
}
