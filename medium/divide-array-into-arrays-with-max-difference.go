package medium

import "sort"

// link: https://leetcode.cn/problems/divide-array-into-arrays-with-maximum-difference/
// thought: 贪心，排序后每3个一组，判断最大值与最小值的差是否超过k
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		res = append(res, nums[i:i+3])
	}
	return res
}
