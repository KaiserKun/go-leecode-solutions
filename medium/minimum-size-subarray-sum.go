package medium

// link: https://leetcode.cn/problems/minimum-size-subarray-sum/
// thought: 滑动窗口，左指针扩大窗口，右指针缩小窗口
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    ans, left, right := n+1, n-1, n-1
    sum := 0
    for left >= 0 {
        sum += nums[left]
        for sum >= target {
            ans = min(ans, right-left+1)
            sum -= nums[right]
            right--
        }
        left--
    }
    if ans > n {
        return 0
    }
    return ans
}