package easy

// link: https://leetcode.cn/problems/final-value-of-variable-after-performing-operations/
// thought: 直接模拟操作过程，遍历操作数组，根据操作类型更新变量值
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, v := range operations {
		if v == "X++" || v == "++X" {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
