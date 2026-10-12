package t_078_617_merge_two_binary_trees

import (
	"reflect"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	var (
		root1     = []int{1, 3, 2, 5}
		root2     = []int{2, 1, 3, 0, 4, 0, 7}
		l         = len(root1)
		queue     = make([]*TreeNode, l)
		expected  = []int{3, 4, 5, 4, 5, 7}
		actual    []int
		traversal func(root *TreeNode)
	)
	for index, val := range root1 {
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
	root1Node := queue[0]
	l = len(root2)
	queue = make([]*TreeNode, l)
	for index, val := range root2 {
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
	root2Node := queue[0]
	root := mergeTrees(root1Node, root2Node)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		actual = append(actual, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
