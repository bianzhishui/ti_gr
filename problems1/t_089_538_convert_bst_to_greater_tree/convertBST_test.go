package t_089_538_convert_bst_to_greater_tree

import (
	"reflect"
	"testing"
)

func TestConvertBST(t *testing.T) {
	var (
		root     = []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{30, 36, 36, 35, 33, 21, 26, 15, 8}
		actual   []int
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
	node := convertBST1(queue[0])
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		actual = append(actual, cur.Val)
		traversal(cur.Left)
		traversal(cur.Right)
	}
	traversal(node)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
