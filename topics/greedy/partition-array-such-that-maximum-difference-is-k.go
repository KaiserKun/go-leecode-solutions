package greedy

import "sort"

// link: https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k/
// thought: 贪心，先排序，然后遍历数组，遇到差值大于k时，分组数加一，并更新当前组的最小值
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	mn := nums[0]
	res := 0
	for _, v := range nums {
		if v-mn > k {
			res++
			mn = v
		}
	}
	return res + 1
}
