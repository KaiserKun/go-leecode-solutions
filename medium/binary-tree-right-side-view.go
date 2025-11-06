package medium

// link: https://leetcode.cn/problems/binary-tree-right-side-view/
// thought: 层序遍历实现，
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	ans := []int{}
	for len(q) > 0 {
		size := len(q)
		// 跳过这一层前面的节点
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i == size-1 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}
