package medium

// link: https://leetcode.cn/problems/subarray-product-less-than-k/
// thought: 左右指针 + 前缀和
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	n := len(nums)
	pro, left, right := 1, 0, 0
	ans := 0
	for right < n {
		pro *= nums[right]
		for pro >= k {
			pro /= nums[left]
			left++
		}
		ans += right - left + 1
		right++
	}
	return ans
}
