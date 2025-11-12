package hard

// link: https://leetcode.cn/problems/shortest-path-visiting-all-nodes/
// thought: 重点是状态，1、当前所处的节点序号；2、跑到这个节点时，哪些节点已经访问；3、走了多远；于是有 u，mask，dist
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type node struct {
		u, mask, dist int
	}
	// seen 记录哪些状态已访问
	seen := make([][]bool, n)
	q := []node{}
	for i := range graph {
		seen[i] = make([]bool, 1<<n-1)
		seen[i][1<<i] = true
		q = append(q, node{i, 1 << i, 0})
	}

	for {
		cur := q[0]
		q = q[1:]
		if cur.mask == 1<<n-1 {
			return cur.dist
		}
		for _, v := range graph[cur.u] {
			maskV := cur.mask | 1<<v
			if seen[v][maskV] {
				continue
			}
			seen[cur.u][1<<v] = true
			q = append(q, node{v, maskV, cur.dist + 1})
		}
	}
}
