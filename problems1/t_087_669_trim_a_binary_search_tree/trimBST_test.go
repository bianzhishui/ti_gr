package t_087_669_trim_a_binary_search_tree

import (
	"container/list"
	"testing"
)

func TestTrimBST(t *testing.T) {
	var (
		root   = []int{1, 0, 2}
		low    = 1
		high   = 2
		l      = len(root)
		queue  = make([]*TreeNode, l)
		actual *TreeNode
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
	actual = trimBST1(queue[0], low, high)
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
	// 验证返回的树是否符合题意
	valValid := false
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		if cur.Val > high || cur.Val < low {
			valValid = true
			return
		}
		if !valValid {
			traversal(cur.Left)
		}
		if !valValid {
			traversal(cur.Right)
		}
	}
	traversal(actual)
	if !valValid {
		t.Logf("%v in [low,high].", actual)
	} else {
		t.Fatalf("%v not in [low,high].", actual)
	}
}
