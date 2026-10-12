package t_083_236_lowest_common_ancestor_of_a_binary_tree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	var (
		root     = []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = 3
		actual   int
	)
	for index, val := range root {
		if val != -1 {
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
	node := lowestCommonAncestor1(queue[0], queue[1], queue[2])
	actual = node.Val
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
