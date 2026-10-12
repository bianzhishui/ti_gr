package t_017_707_design_linked_list

import (
	"reflect"
	"testing"
)

func testData1(td []int) (linkedList MyLinkedList1) {
	var (
		virhead    = &ListNode1{Next: nil}
		virheadTmp = virhead
		l          int
	)
	l = len(td)
	if l == 0 {
		linkedList = MyLinkedList1{Header: nil, Tail: nil, Lens: l}
		return
	}
	for i := 0; i < l; i++ {
		virheadTmp.Next = &ListNode1{Val: td[i], Next: nil}
		virheadTmp = virheadTmp.Next
	}
	linkedList = MyLinkedList1{Header: virhead.Next, Tail: virheadTmp, Lens: l}
	return
}

func TestMyLinkedList1_Constructor1(t *testing.T) {
	var (
		expected = MyLinkedList1{Header: nil, Tail: nil, Lens: 0}
		actual   MyLinkedList1
	)
	actual = Constructor1()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestMyLinkedList1_GetMyLinkedListLen1(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		expected = 3
		actual   int
	)
	linkedList := testData1(td)
	actual = linkedList.GetMyLinkedListLen()
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_PrintMyLinkedList(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		expected = []int{0, 1, 2}
		actual   []int
	)
	linkedList := testData1(td)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_AddAtTail(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		val      = 6
		expected = []int{0, 1, 2, 6}
		actual   []int
	)

	linkedList := testData1(td)
	linkedList.AddAtTail(val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_AddAtHead(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		val      = 6
		expected = []int{6, 0, 1, 2}
		actual   []int
	)

	linkedList := testData1(td)
	linkedList.AddAtHead(val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_AddAtIndex(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		val      = 6
		expected = []int{0, 1, 6, 2}
		actual   []int
	)

	linkedList := testData1(td)
	linkedList.AddAtIndex(index, val)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_DeleteAtIndex(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		expected = []int{0, 1}
		actual   []int
	)

	linkedList := testData1(td)
	linkedList.DeleteAtIndex(index)
	actual = linkedList.PrintMyLinkedList()
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestMyLinkedList1_Get(t *testing.T) {
	var (
		td       = []int{0, 1, 2}
		index    = 2
		expected = 2
		actual   int
	)

	linkedList := testData1(td)
	actual = linkedList.Get(index)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
