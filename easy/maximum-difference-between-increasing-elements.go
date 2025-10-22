package easy

// link: https://leetcode.com/problems/maximum-difference-between-increasing-elements/
// thought: 遍历时检查值的大小，小于当前最小值，则改变min，大于则判断是否大于当前差值
func maximumDifference(nums []int) int {
	n := len(nums)
	res := -1
	min := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > min {
			if nums[i]-min > res {
				res = nums[i] - min
			}
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	return res
}
