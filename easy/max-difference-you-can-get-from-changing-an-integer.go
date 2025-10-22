package easy

import "strconv"

// link: https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
// thought: 计算把某一非 9 的高位全部替换为 9 后的最大数；以及把某一位替换为 0 后的最小数，返回两者之差。
func minMaxDifference(num int) int {
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

	maxStr := s
	for i := 0; i < n; i++ {
		if s[i] != '9' {
			maxStr = replaceAll(s[i], '9')
			break
		}
	}

	minStr := s
	// 找到第一个可以替换为 '0' 的高位
	for i := 0; i < n; i++ {
		if s[i] != '0' {
			minStr = replaceAll(s[i], '0')
			break
		}
	}

	maxV, _ := strconv.Atoi(maxStr)
	minV, _ := strconv.Atoi(minStr)
	return maxV - minV
}
