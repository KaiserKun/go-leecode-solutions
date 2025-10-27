package greedy

import "sort"

// link: https://leetcode.cn/problems/divide-array-into-arrays-with-max-difference/
// thought: 贪心算法，先排序，然后每三个一组检查最大差值是否满足要求
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
