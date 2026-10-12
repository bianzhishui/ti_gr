package t_018_206_reverse_linked_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var (
		headList = []int{1, 2, 3, 4, 5}
		head     = &ListNode{}
		headTmp  = head
		expected = []int{5, 4, 3, 2, 1}
		actual   []int
	)
	for _, v := range headList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp = reverseList(head.Next)
	actual = make([]int, 0, len(headList))
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
