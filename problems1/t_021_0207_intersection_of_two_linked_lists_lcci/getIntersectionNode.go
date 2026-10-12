package t_021_0207_intersection_of_two_linked_lists_lcci

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	// 移动较长链表与短链表对齐
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 字典
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	dic := make(map[*ListNode]bool)
	for headA != nil {
		dic[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		_, ok := dic[headB]
		if ok {
			return headB
		} else {
			headB = headB.Next
		}
	}
	return nil
}
