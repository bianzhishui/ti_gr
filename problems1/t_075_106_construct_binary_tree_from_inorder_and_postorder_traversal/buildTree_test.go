package t_075_106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var (
		inorder   = []int{9, 3, 15, 20, 7}
		postorder = []int{9, 15, 7, 20, 3}
		// 取前序遍历值的顺序校验
		expected  = []int{3, 9, 20, 15, 7}
		actual    = []int{}
		traversal func(root *TreeNode)
	)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		actual = append(actual, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	root := buildTree(inorder, postorder)
	traversal(root)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
