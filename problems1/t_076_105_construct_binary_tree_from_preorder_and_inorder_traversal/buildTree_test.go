package t_076_105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var (
		inorder   = []int{9, 3, 15, 20, 7}
		postorder = []int{3, 9, 20, 15, 7}
		// 取后序遍历值的顺序校验
		expected  = []int{9, 15, 7, 20, 3}
		actual    = []int{}
		traversal func(root *TreeNode)
	)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		actual = append(actual, root.Val)
	}
	root := buildTree(postorder, inorder)
	traversal(root)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
