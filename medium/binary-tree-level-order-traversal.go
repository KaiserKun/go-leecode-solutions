package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// link: https://leetcode.cn/problems/binary-tree-level-order-traversal/
// thought: 层序遍历的标准实现，bfs
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := []int{}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ans = append(ans, level)
	}
	return ans
}
