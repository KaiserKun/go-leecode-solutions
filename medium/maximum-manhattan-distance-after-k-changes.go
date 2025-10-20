package medium

// link: https://leetcode.cn/problems/maximum-manhattan-distance-after-k-changes/
// thought: 贪心，每次改变一个字符可以使曼哈顿距离增加2，计算当前位置的曼哈顿距离加上k*2与当前位置i+1的最小值
func maxDistance(s string, k int) int {
	x := 0
	y := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		ans = max(ans, min(abs(x)+abs(y)+k*2, i+1))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
