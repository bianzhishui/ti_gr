package t_066_572_subtree_of_another_tree

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	var (
		root     = []int{1, 2, 2, 3, 4, 4, 3}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = true
		actual   bool
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
	actual = isSubtree(queue[0], queue[0])
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
