package t_056_515_find_largest_value_in_each_tree_row

import (
	"reflect"
	"testing"
)

func TestLargestValues(t *testing.T) {
	var (
		root     = []int{1, 3, 2, 5, 3, 0, 9}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{1, 3, 9}
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
	actual = largestValues(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
