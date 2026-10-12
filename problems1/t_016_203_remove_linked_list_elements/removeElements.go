package t_016_203_remove_linked_list_elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	// 定义虚拟头
	virHead := ListNode{Next: head}
	// 定义临时变量用来遍历列表 循环条件为链表下一个值不为空
	for headTmp := &virHead; headTmp.Next != nil; {
		// 判断下一个链表节点的值等于 val 若等于 val 则删除 若不等于 val 则向下遍历
		if headTmp.Next.Val == val {
			headTmp.Next = headTmp.Next.Next
		} else {
			headTmp = headTmp.Next
		}
	}
	return virHead.Next
}

// 递归
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func removeElements2(head *ListNode, val int) *ListNode {
	// 递归返回条件为当前链表节点为 nil
	if head == nil {
		return head
	}
	// 递归到最后
	head.Next = removeElements(head.Next, val)
	// 从最后判断链表节点值是否为 val 若是则返回当前节点的 Next 否则返回链表节点
	if head.Val == val {
		return head.Next
	}
	return head
}
