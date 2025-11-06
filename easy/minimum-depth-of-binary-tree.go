package easy

import "container/list"

// link: https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// thought: 标准二叉树实现，层级遍历时，判断是否为子节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	ans := 0
	for q.Len() > 0 {
		size := q.Len()
		ans++
		for i := 0; i < size; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return ans
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}
	return ans
}
