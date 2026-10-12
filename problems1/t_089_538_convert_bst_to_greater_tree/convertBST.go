package t_089_538_convert_bst_to_greater_tree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func convertBST(root *TreeNode) *TreeNode {
	var traversal func(cur *TreeNode)
	sum := 0
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Right)
		cur.Val += sum
		sum = cur.Val
		traversal(cur.Left)
	}
	traversal(root)
	return root
}

// 迭代
func convertBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	sum := 0
	cur := root
	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Right
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			cur.Val += sum
			sum = cur.Val
			cur = cur.Left
		}
	}
	return root
}
