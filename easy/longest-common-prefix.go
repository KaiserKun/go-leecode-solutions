package easy

// link: https://leetcode.cn/problems/longest-common-prefix/
// thought: 逐次比较字符串数组中的元素
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i < n; i++ {
		prefix = compare(prefix, strs[i])
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func compare(str1, str2 string) string {
	n := min(len(str1), len(str2))
	i := 0
	for i < n && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}
