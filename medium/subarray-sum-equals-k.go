package medium

// link: https://leetcode.cn/problems/subarray-sum-equals-k/
// thought: 前缀和 + 哈希表
func subarraySum(nums []int, k int) int {
	cnt, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			cnt += mp[pre-k]
		}
		mp[pre] += 1
	}
	return cnt
}
