package t_018_206_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseList(head *ListNode) *ListNode {
	// 定义新的链表头节点
	var prev *ListNode
	curr := head
	for curr != nil {
		// 临时变量 指向改变前下一个节点
		next := curr.Next
		// 改变下一个节点
		curr.Next = prev
		// 更换头节点
		prev = curr
		// 指向改变前的下一个节点
		curr = next
	}
	return prev
}

// 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
