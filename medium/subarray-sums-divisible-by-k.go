package medium

// link: https://leetcode.cn/problems/subarray-sums-divisible-by-k/
// thought: 同余，表示中间的子数组之和可以整除k
func subarraysDivByK(nums []int, k int) int {
	modcnt := map[int]int{0: 1}
	prefix := 0
	ans := 0

	for _, num := range nums {
		prefix += num

		mod := prefix % k
		if mod < 0 {
			mod += k
		}
		ans += modcnt[mod]
		modcnt[mod]++
	}
	return ans
}
