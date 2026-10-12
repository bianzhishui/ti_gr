package t_021_0207_intersection_of_two_linked_lists_lcci

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	var (
		headList   = []int{8, 4, 5}
		headAList  = []int{4, 1}
		headBList  = []int{5, 0, 1}
		headA      = &ListNode{}
		headB      = &ListNode{}
		headPublic = &ListNode{}
		headTmp    = headPublic
		expected   = headList
		actual     []int
	)
	for _, v := range headList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp = headA
	for _, v := range headAList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp.Next = headPublic.Next
	headTmp = headB
	for _, v := range headBList {
		headTmp.Next = &ListNode{Val: v}
		headTmp = headTmp.Next
	}
	headTmp.Next = headPublic.Next
	headTmp = getIntersectionNode(headA.Next, headB.Next)
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
