package medium

// link: https://leetcode.com/problems/zigzag-conversion/
// thought: 利用行和列的关系，将字符串按Z字形排列
func convert(s string, numRows int) string {
	// 行数为1时，直接返回原字符串
	if numRows == 1 {
		return s
	}
	// 利用行数和列数的关系，将字符串按Z字形排列
	rows := make([][]byte, numRows)
	r, flag := 0, 1
	// 利用flag上下移动
	for i := 0; i < len(s); i++ {
		rows[r] = append(rows[r], s[i])
		r += flag
		if r == 0 || r == numRows-1 {
			flag = -flag
		}
	}
	// 将所有行拼接成一个
	var ret []byte
	for _, row := range rows {
		ret = append(ret, row...)
	}
	return string(ret)
}
