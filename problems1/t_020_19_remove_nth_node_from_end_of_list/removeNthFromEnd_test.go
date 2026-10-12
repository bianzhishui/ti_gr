package t_020_19_remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var (
		headList = []int{1, 2, 3, 4, 5}
		head     = &ListNode{}
		headTmp  = head
		n        = 2
		expected = []int{1, 2, 3, 5}
		actual   []int
	)

	for _, v := range headList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp = removeNthFromEnd(head.Next, n)
	for ; headTmp != nil; headTmp = headTmp.Next {
		actual = append(actual, headTmp.Val)
	}
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
