package t_079_700_search_in_a_binary_search_tree

import (
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	var (
		root      = []int{4, 2, 7, 1, 3}
		val       = 2
		l         = len(root)
		queue     = make([]*TreeNode, l)
		expected  = []int{2, 1, 3}
		actual    []int
		traversal func(node *TreeNode)
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
	node := searchBST2(queue[0], val)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		actual = append(actual, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(node)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
