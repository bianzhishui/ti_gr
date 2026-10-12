package t_077_654_maximum_binary_tree

import (
	"reflect"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	var (
		nums = []int{3, 2, 1, 6, 0, 5}
		// 取前序遍历值的顺序校验
		expected  = []int{6, 3, 2, 1, 5, 0}
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
	root := constructMaximumBinaryTree(nums)
	traversal(root)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
