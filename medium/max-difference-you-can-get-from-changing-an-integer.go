package medium

import (
	"strconv"
)

// link: https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// 优化思路：避免使用 strings.ReplaceAll 多次分配，直接基于字节数组一次性生成替换结果。
// thought：计算把某一非 9 的高位全部替换为 9 后的最大数；以及把某一位替换为 1/0 后的最小数，返回两者之差。
func maxDiff(num int) int {
	s := strconv.Itoa(num)
	n := len(s)

	// 生成替换后的字符串（将 s 中所有等于 from 的字符替换为 to）
	replaceAll := func(from, to byte) string {
		buf := make([]byte, n)
		for i := 0; i < n; i++ {
			if s[i] == from {
				buf[i] = to
			} else {
				buf[i] = s[i]
			}
		}
		return string(buf)
	}

	// 计算最大值：找到第一个不是 '9' 的位，把该位的所有相同数字替换为 '9'
	maxStr := s
	for i := 0; i < n; i++ {
		if s[i] != '9' {
			maxStr = replaceAll(s[i], '9')
			break
		}
	}

	// 计算最小值：分两种情况
	// 1) 如果最高位不是 '1'，把最高位的所有相同数字替换为 '1'
	// 2) 否则，找到第一个既不是最高位字符也不是 '0' 的位，把它的所有相同数字替换为 '0'
	minStr := s
	if s[0] != '1' {
		minStr = replaceAll(s[0], '1')
	} else {
		// 找到第一个可以替换为 '0' 的高位
		for i := 1; i < n; i++ {
			if s[i] != '0' && s[i] != s[0] {
				minStr = replaceAll(s[i], '0')
				break
			}
		}
	}

	maxV, _ := strconv.Atoi(maxStr)
	minV, _ := strconv.Atoi(minStr)
	return maxV - minV
}
