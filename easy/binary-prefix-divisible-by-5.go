package easy

// link: https://leetcode.cn/problems/binary-prefix-divisible-by-5/
// thought: 每次判断即可，注意sum只需要保留余数即可
func prefixesDivBy5(nums []int) []bool {
	sum := 0
	ans := make([]bool, len(nums))
	for i, num := range nums {
		sum = (sum*2 + num) % 5
		if sum == 0 {
			ans[i] = true
		}
	}
	return ans
}
