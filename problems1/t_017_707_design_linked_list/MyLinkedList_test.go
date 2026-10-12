package t_017_707_design_linked_list

import (
	"reflect"
	"testing"
)

func testData(td []int) (linkedList MyLinkedList) {
	var (
		head    = &ListNode{Next: nil}
		headTmp = head
	)
	for i := 0; i < len(td); i++ {
		headTmp.Next = &ListNode{Val: td[i], Next: nil}
		headTmp = headTmp.Next
	}
	linkedList = MyLinkedList{Node: head}
	return
}

func TestMyLinkedList_Constructor(t *testing.T) {
	var (
		virHead  = ListNode{Next: nil}
		expected = MyLinkedList{Node: &virHead}
		actual   MyLinkedList
	)
	actual = Constructor()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestMyLinkedList_GetMyLinkedListLen(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		expected = 3
		actual   int
	)
	linkedList := testData(td)
	actual = linkedList.GetMyLinkedListLen()
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_PrintMyLinkedList(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		expected = []int{0, 1, 2}
		actual   []int
	)
	linkedList := testData(td)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		val      = 6
		expected = []int{0, 1, 2, 6}
		actual   []int
	)

	linkedList := testData(td)
	linkedList.AddAtTail(val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		val      = 6
		expected = []int{6, 0, 1, 2}
		actual   []int
	)

	linkedList := testData(td)
	linkedList.AddAtHead(val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		val      = 6
		expected = []int{0, 1, 6, 2}
		actual   []int
	)

	linkedList := testData(td)
	linkedList.AddAtIndex(index, val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		expected = []int{0, 1}
		actual   []int
	)

	linkedList := testData(td)
	linkedList.DeleteAtIndex(index)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList_Get(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		expected = 2
		actual   int
	)

	linkedList := testData(td)
	actual = linkedList.Get(index)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
