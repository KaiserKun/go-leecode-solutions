package easy

// link: https://leetcode.cn/problems/palindrome-number/
// thought: 反转数字然后比较
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp, cmp := x, 0
	for tmp != 0 {
		cmp = cmp*10 + tmp%10
		tmp /= 10
	}
	return cmp == x
}
