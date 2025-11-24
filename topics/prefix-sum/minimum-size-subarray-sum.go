package ps

// link: https://leetcode.cn/problems/minimum-size-subarray-sum/
// thought: 利用sum表示前缀和，滑动窗口解题
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, left, right := n+1, 0, 0
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans > n {
		return 0
	}
	return ans
}
