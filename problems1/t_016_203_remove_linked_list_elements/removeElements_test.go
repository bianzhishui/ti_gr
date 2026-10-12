package t_016_203_remove_linked_list_elements

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	var (
		headList = []int{1, 2, 6, 3, 4, 5, 6}
		head     = &ListNode{Next: nil}
		headTmp  = head
		val      = 6
		expected = []int{1, 2, 3, 4, 5}
		actual   []int
	)

	for i := 0; i < len(headList); i++ {
		headTmp.Next = &ListNode{Val: headList[i], Next: nil}
		headTmp = headTmp.Next
	}

	head = removeElements(head.Next, val)
	for ; head != nil; head = head.Next {
		actual = append(actual, head.Val)
	}

	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
