package easy

// link: https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
// thought: 使用一个数组记录每个位置的连续递增对数，然后检查是否存在满足条件的子数组
func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if n < 2*k {
		return false
	}
	// inc[i] 表示以 i 结尾的连续递增对数
	inc := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			inc[i] = inc[i-1] + 1
		} else {
			inc[i] = 0
		}
	}
	// 检查inc数组，判断是否存在满足条件的子数组
	for i := 0; i <= n-2*k; i++ {
		if inc[i+k-1] >= k-1 && inc[i+2*k-1] >= k-1 {
			return true
		}
	}
	return false
}
