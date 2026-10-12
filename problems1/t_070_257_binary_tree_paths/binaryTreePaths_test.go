package t_070_257_binary_tree_paths

import (
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	var (
		root     = []int{1, 2, 3, 0, 5}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []string{"1->2->5", "1->3"}
		actual   []string
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
	actual = binaryTreePaths(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
