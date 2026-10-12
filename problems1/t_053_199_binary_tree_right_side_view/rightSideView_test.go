package t_053_199_binary_tree_right_side_view

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	var (
		root     = []int{3, 9, 20, 0, 0, 15, 7}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{3, 20, 7}
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
	actual = rightSideView(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
