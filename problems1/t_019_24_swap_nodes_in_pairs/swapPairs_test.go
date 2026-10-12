package t_019_24_swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var (
		headList = []int{1, 2, 3, 4}
		head     = &ListNode{}
		headTmp  = head
		expected = []int{2, 1, 4, 3}
		actual   []int
	)

	for _, v := range headList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp = swapPairs(head.Next)
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
