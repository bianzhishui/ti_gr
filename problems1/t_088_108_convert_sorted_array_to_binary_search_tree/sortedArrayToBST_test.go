package t_088_108_convert_sorted_array_to_binary_search_tree

import (
	"container/list"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	var (
		nums   = []int{-10, -3, 0, 5, 9}
		actual *TreeNode
	)
	actual = sortedArrayToBST1(nums)
	// 验证是二叉搜索树
	var isValidBST func(cur *TreeNode) bool
	isValidBST = func(cur *TreeNode) bool {
		stack := list.New()
		var pre *TreeNode
		for cur != nil || stack.Len() != 0 {
			if cur != nil {
				stack.PushBack(cur)
				cur = cur.Left
			} else {
				cur = stack.Remove(stack.Back()).(*TreeNode)
				if pre != nil && cur.Val <= pre.Val {
					return false
				}
				pre = cur
				cur = cur.Right
			}
		}
		return true
	}
	if isValidBST(actual) {
		t.Logf("%v is BST.", actual)
	} else {
		t.Fatalf("%v is not BST.", actual)
	}
}
