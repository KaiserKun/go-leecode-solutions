package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// link: https://leetcode.cn/problems/average-of-levels-in-binary-tree/
// thought: 层序遍历的标准实现，计算值即可
func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	ans := []float64{}
	for len(q) > 0 {
		size := len(q)
		sum := 0.0
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum += float64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, sum/float64(size))
	}
	return ans
}
