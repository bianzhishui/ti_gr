package t_068_222_count_complete_tree_nodes

import (
	"testing"
)

func TestCountNodes(t *testing.T) {
	var (
		root     = []int{1, 2, 3, 4, 5, 6}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = 6
		actual   int
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
	actual = countNodes(queue[0])
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
