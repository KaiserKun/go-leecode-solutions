package medium

import "fmt"

// link: https://leetcode.cn/problems/complex-number-multiplication/
// thought: 直接按格式输入到指定的变量，然后按复数乘法公式计算
func complexNumberMultiply(num1 string, num2 string) string {
	// 利用fmt的输入，直接按格式输入指定的变量
	var a, b, c, d int
	fmt.Sscanf(num1, "%d+%di", &a, &b)
	fmt.Sscanf(num2, "%d+%di", &c, &d)
	return fmt.Sprintf("%d+%di", a*c-b*d, a*d+b*c)
}
