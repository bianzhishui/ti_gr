package t_080_98_validate_binary_search_tree

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

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var traversal func(node *TreeNode) bool
	traversal = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := traversal(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		} else {
			pre = node
		}
		right := traversal(node.Right)
		return left && right
	}
	return traversal(root)
}

// 迭代中序遍历
func isValidBST1(root *TreeNode) bool {
	stack := list.New()
	var pre *TreeNode
	for root != nil || stack.Len() != 0 {
		if root != nil {
			stack.PushBack(root)
			root = root.Left
		} else {
			root = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil && root.Val <= pre.Val {
				return false
			}
			pre = root
			root = root.Right
		}
	}
	return true
}
