package t_071_404_sum_of_left_leaves

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
func sumOfLeftLeaves(root *TreeNode) int {
	ret := 0
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 判断是否是左叶子节点 若是则将值加在ret中
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ret += node.Left.Val
		}
		// 继续遍历左右节点
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return ret
}

// 迭代
func sumOfLeftLeaves1(root *TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}
	// 栈
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ret += node.Left.Val
		}
		if node.Right != nil {

			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ret
}
