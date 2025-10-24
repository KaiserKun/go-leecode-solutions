package easy

// link: https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// thought: 计算相邻元素之间的最大差值
func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	res := abs(nums[0] - nums[n-1])
	for i := 1; i < n; i++ {
		res = max(res, abs(nums[i]-nums[i-1]))
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
