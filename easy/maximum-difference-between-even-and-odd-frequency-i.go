package easy

// link: https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
// thought: 计算字符出现奇数次和偶数次的最大差值
func maxDifference(s string) int {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	maxOdd, minEven := 0, 0
	for _, v := range freq {
		// 出现偶数次
		if v%2 == 0 {
			if v < minEven || minEven == 0 {
				minEven = v
			}
		} else {
			if v > maxOdd || maxOdd == 0 {
				maxOdd = v
			}
		}
	}
	return maxOdd - minEven
}
