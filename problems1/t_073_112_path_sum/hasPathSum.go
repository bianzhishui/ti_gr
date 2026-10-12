package t_073_112_path_sum

import (
	"container/list"
)

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

func hasPathSum(root *TreeNode, targetSum int) bool {
	ret := false
	var traversal func(node *TreeNode, sum int)
	traversal = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			ret = true
			return
		}
		if !ret {
			traversal(node.Left, sum)
		}
		if !ret {
			traversal(node.Right, sum)
		}
	}
	traversal(root, 0)
	return ret
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	var traversal func(node *TreeNode, sum int) bool
	traversal = func(node *TreeNode, sum int) bool {
		// if node == nil {
		// 	return false
		// }
		// // 叶子节点且sum为0 说明找到了符合题意的路径
		// if node.Left == nil && node.Right == nil && sum == node.Val {
		// 	return true
		// }
		// if traversal(node.Left, sum-node.Val) {
		// 	return true
		// }
		// if traversal(node.Right, sum-node.Val) {
		// 	return true
		// }
		// return false
		return node != nil && ((node.Left == nil && node.Right == nil && sum == node.Val) || traversal(node.Left, sum-node.Val) || traversal(node.Right, sum-node.Val))
	}
	return traversal(root, targetSum)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	stack.PushBack([]interface{}{root, targetSum})
	for stack.Len() != 0 {
		nodeInfo := stack.Remove(stack.Back()).([]interface{})
		node := nodeInfo[0].(*TreeNode)
		sum := nodeInfo[1].(int)
		if node.Left == nil && node.Right == nil && node.Val == sum {
			return true
		}
		if node.Right != nil {
			stack.PushBack([]interface{}{node.Right, sum - node.Val})
		}
		if node.Left != nil {
			stack.PushBack([]interface{}{node.Left, sum - node.Val})
		}
	}
	return false
}
