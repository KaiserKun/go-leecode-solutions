package medium

import "sort"

// link: https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k/
// thought: 贪心，排序后从左到右遍历，遇到大于k的差值时分组
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	mn := nums[0]
	res := 0
	for _, x := range nums {
		if x-mn > k {
			res++
			mn = x
		}
	}
	return res + 1
}
