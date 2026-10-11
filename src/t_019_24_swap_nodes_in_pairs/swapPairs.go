package t_019_24_swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	tempHead := dummyHead
	for tempHead.Next != nil && tempHead.Next.Next != nil {
		node1 := tempHead.Next
		node2 := tempHead.Next.Next
		node1.Next = node2.Next
		node2.Next = node1
		tempHead.Next = node2
		tempHead = node1
	}
	return dummyHead.Next
}

// 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
