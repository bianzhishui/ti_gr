package t_084_235_lowest_common_ancestor_of_a_binary_search_tree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	var (
		root     = []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = 6
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
	node := lowestCommonAncestor(queue[0], queue[1], queue[2])
	actual = node.Val
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
