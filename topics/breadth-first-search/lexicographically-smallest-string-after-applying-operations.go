package breadthfirstsearch

// link: https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/
// thought: 分别处理b为偶数和奇数的情况,b为偶数时只能改变偶数位，b为奇数时可以改变所有位； 使用集合记录所有可能的字符串，最后返回字典序最小的字符串
//
//	使用切片记录所有可能出现的字符串，最后通过遍历切片找到字典序最小的字符串
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	visited := make(map[string]bool)
	queue := []string{s}
	visited[s] = true
	minStr := s
	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		if current < minStr {
			minStr = current
		}

		// 操作1：对所有奇数位加a（更高效的策略）
		bytes := []byte(current)
		for i := 1; i < n; i += 2 {
			digit := int(bytes[i] - '0')
			newDigit := (digit + a) % 10
			bytes[i] = byte(newDigit) + '0'
		}
		newStr1 := string(bytes)
		if !visited[newStr1] {
			visited[newStr1] = true
			queue = append(queue, newStr1)
		}

		// 操作2：旋转
		rotated := current[n-b:] + current[:n-b]
		if !visited[rotated] {
			visited[rotated] = true
			queue = append(queue, rotated)
		}
	}

	return minStr
}
