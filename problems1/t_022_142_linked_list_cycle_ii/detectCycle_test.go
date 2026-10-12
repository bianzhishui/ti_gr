package t_022_142_linked_list_cycle_ii

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// 定义环形链表
	// 3 2 0 4 2 0 4 2 0 4 ...
	var (
		headList   = []int{3, 2, 0, 4}
		headPublic = &ListNode{}
		headTmp    = headPublic
		expected   = 2
		actual     int
	)
	for _, v := range headList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	// 4 节点 指向 2 节点
	headTmp.Next = headPublic.Next.Next
	headTmp = detectCycle(headPublic.Next)
	actual = headTmp.Val
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
