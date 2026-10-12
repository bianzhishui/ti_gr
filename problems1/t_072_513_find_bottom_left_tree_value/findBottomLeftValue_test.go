package t_072_513_find_bottom_left_tree_value

import (
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	var (
		root     = []int{1, 2, 3, 4, 0, 5, 6, 0, 0, 0, 0, 7}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = 7
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
	actual = findBottomLeftValue1(queue[0])
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
