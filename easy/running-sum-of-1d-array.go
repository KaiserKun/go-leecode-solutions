package easy

// link: https://leetcode.cn/problems/running-sum-of-1d-array/
// thought: 直接在原数组修改即可
func runningSum(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
