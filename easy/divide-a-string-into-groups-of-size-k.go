package easy

import "strings"

// link: https://leetcode.cn/problems/divide-a-string-into-groups-of-size-k/
// thought: 将字符串分割成固定大小的块，不足的块用 fill 填充
func divideString(s string, k int, fill byte) []string {
	n := len(s)
	var res []string
	cur := 0
	for cur < n {
		end := cur + k
		if end > n {
			end = n
		}
		res = append(res, s[cur:end])
		cur += k
	}
	if cur != n {
		res[len(res)-1] += strings.Repeat(string(fill), cur-n)
	}
	return res
}
