package medium

import "math"

// link: https://leetcode.cn/problems/string-to-integer-atoi/
// thought: 模拟atoi函数的行为，处理空格、符号、数字和溢出情况
func myAtoi(s string) int {
	n := len(s)
	i := 0
	flag := 1
	ans := 0
	// 跳过空格
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			flag = -1
		}
		i++
	}
	const min_div = math.MaxInt32 / 10
	const min_mod = math.MaxInt32 % 10
	for i < n && isdigit(s[i]) {
		digit := int(s[i] - '0')
		if ans > min_div || (ans == min_div && digit > min_mod) {
			if flag == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		ans = ans*10 + digit
		i++
	}
	return ans * flag
}

func isdigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
