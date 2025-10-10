package medium

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// link: https://leetcode.cn/problems/add-two-numbers/
// thought: 模拟加法过程，使用一个变量 carry 来处理进位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 使用哑节点简化结果链表的构造
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		// 防止访问空指针
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 创建新节点存储当前位的结果
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}

	return dummy.Next
}
