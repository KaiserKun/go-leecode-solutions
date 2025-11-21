package ps

// link: https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
// thought: 标准前缀和，求正方形的和，用二分查找确定最大边长
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])

	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefix[i][j] = prefix[i][j-1] + prefix[i-1][j] - prefix[i-1][j-1] + mat[i-1][j-1]
		}
	}

	left, right := 1, min(m, n)
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		found := false
		for i := mid; i <= m; i++ {
			for j := mid; j <= n; j++ {
				sum := prefix[i][j] - prefix[i-mid][j] - prefix[i][j-mid] + prefix[i-mid][j-mid]
				if sum <= threshold {
					found = true
					ans = max(ans, mid)
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
