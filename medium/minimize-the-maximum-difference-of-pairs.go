package medium

import "sort"

// link: https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/
// thought: 排序后使用二分查找最大差值的最小值，检查函数贪心地配对
func minimizeMax(nums []int, p int) int {
	n := len(nums)
	// 排序
	sort.Ints(nums)
	// 检查最大差值为mid时，是否能组成p对
	check := func(mid int) bool {
		count := 0
		i := 1
		for i < n {
			if nums[i]-nums[i-1] <= mid {
				count++
				i += 2
			} else {
				i++
			}
		}
		return count >= p
	}
	// 二分查找最小的最大差值
	lo, hi := 0, nums[n-1]-nums[0]
	ans := hi
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(mid) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
