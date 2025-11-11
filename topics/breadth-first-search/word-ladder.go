package bfs

// link: https://leetcode.cn/problems/word-ladder/
// thought: 使用BFS，状态是现在的string
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	ans := 1
//	visited := make(map[string]bool)
//	q := []string{beginWord}
//	visited[beginWord] = true
//
//	for len(q) > 0 {
//		size := len(q)
//		ans++
//		for i := 0; i < size; i++ {
//			tmp := q[0]
//			q = q[1:]
//			for _, word := range wordList {
//				if findDif(tmp, word) == 1 && !visited[word] {
//					if findDif(word, endWord) == 0 {
//						return ans
//					}
//					visited[word] = true
//					q = append(q, word)
//				}
//			}
//		}
//	}
//	return 0
//}
//
//func findDif(src, dst string) int {
//	n := len(src)
//	cnt := 0
//	for i := 0; i < n; i++ {
//		if src[i] != dst[i] {
//			cnt++
//		}
//	}
//	return cnt
//}

// 双向BFS实现
func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	if beginWord == endWord {
		return 1
	}

	beginQueue := []string{beginWord}
	endQueue := []string{endWord}
	beginVisited := make(map[string]int)
	endVisited := make(map[string]int)
	beginVisited[beginWord] = 1
	endVisited[endWord] = 1

	for len(beginQueue) > 0 && len(endQueue) > 0 {
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
			beginVisited, endVisited = endVisited, beginVisited
		}
		size := len(beginQueue)
		for i := 0; i < size; i++ {
			cur := beginQueue[0]
			beginQueue = beginQueue[1:]
			for j := 0; j < len(cur); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == cur[j] {
						continue
					}

					newWord := cur[:j] + string(c) + cur[j+1:]

					if steps, ok := endVisited[newWord]; ok {
						return beginVisited[cur] + steps
					}
					if wordSet[newWord] && beginVisited[newWord] == 0 {
						beginVisited[newWord] = beginVisited[cur] + 1
						beginQueue = append(beginQueue, newWord)
					}

				}
			}

		}
	}
	return 0
}
