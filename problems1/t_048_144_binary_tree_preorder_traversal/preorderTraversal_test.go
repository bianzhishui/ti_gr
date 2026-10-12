package t_048_144_binary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var (
		root     = []int{1, 0, 2, 0, 0, 3}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{1, 2, 3}
		actual   []int
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &TreeNode{Val: val}
		}
	}
	for index, val := range queue {
		if val == nil {
			continue
		}
		i := index * 2
		if i+1 < l && queue[i+1] != nil {
			val.Left = queue[i+1]
		}
		if i+2 < l && queue[i+2] != nil {
			val.Right = queue[i+2]
		}
	}
	actual = preorderTraversal(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestPreorderTraversal1(t *testing.T) {
	var (
		root     = []int{1, 0, 2, 0, 0, 3}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{1, 2, 3}
		actual   []int
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &TreeNode{Val: val}
		}
	}
	for index, val := range queue {
		if val == nil {
			continue
		}
		i := index * 2
		if i+1 < l && queue[i+1] != nil {
			val.Left = queue[i+1]
		}
		if i+2 < l && queue[i+2] != nil {
			val.Right = queue[i+2]
		}
	}
	actual = preorderTraversal1(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestPreorderTraversal2(t *testing.T) {
	var (
		root     = []int{1, 0, 2, 0, 0, 3}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{1, 2, 3}
		actual   []int
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &TreeNode{Val: val}
		}
	}
	for index, val := range queue {
		if val == nil {
			continue
		}
		i := index * 2
		if i+1 < l && queue[i+1] != nil {
			val.Left = queue[i+1]
		}
		if i+2 < l && queue[i+2] != nil {
			val.Right = queue[i+2]
		}
	}
	actual = preorderTraversal2(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
