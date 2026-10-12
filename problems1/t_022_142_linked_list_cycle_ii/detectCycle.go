package t_022_142_linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func detectCycle(head *ListNode) *ListNode {
	// 定义快慢指针
	slow, fast := head, head
	// 当快指针为空或快指针下一个节点为空 说明无环
	for fast != nil && fast.Next != nil {
		// 慢指针
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

// 哈希表
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func detectCycle2(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
