package hard

// link: https://leetcode.cn/problems/sum-of-k-mirror-numbers/
// thought: 按升序生成十进制回文数，对每个回文数判断其在 k 进制下是否也是回文，累加前 n 个符合条件的十进制值
func kMirror(k int, n int) int64 {
	digit := make([]int, 100)
	left, count, ans := 1, 0, int64(0)
	for count < n {
		right := left * 10
		// op = 0 表示枚举奇数长度回文，op = 1 表示枚举偶数长度回文
		for op := 0; op < 2; op++ {
			// 枚举 i'
			for i := left; i < right && count < n; i++ {
				combined := int64(i)
				x := i
				// 去掉最中间的数字（如果是偶数长度回文则不去掉）
				if op == 0 {
					x = i / 10
				}
				for x > 0 {
					combined = combined*10 + int64(x%10)
					x /= 10
				}
				if isPalindrome(combined, k, digit) {
					count++
					ans += combined
				}
			}
		}
		left = right
	}
	return ans
}

// isPalindrome 判断 x 在 k 进制下是否为回文数
func isPalindrome(x int64, k int, digit []int) bool {
	length := -1
	for x > 0 {
		length++
		digit[length] = int(x % int64(k))
		x /= int64(k)
	}
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		if digit[i] != digit[j] {
			return false
		}
	}
	return true
}
