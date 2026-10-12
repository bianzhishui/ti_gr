package t_072_513_find_bottom_left_tree_value

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

// 迭代
func findBottomLeftValue(root *TreeNode) int {
	ret := 0
	maxDepth := -1
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		// 如果为叶子节点
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				ret = node.Val
				maxDepth = depth
			}
			return
		}
		// 先左后右最后存下来的值为最左值
		if node.Left != nil {
			traversal(node.Left, depth+1)
		}
		if node.Right != nil {
			traversal(node.Right, depth+1)
		}
	}
	traversal(root, 0)
	return ret
}

// 迭代
func findBottomLeftValue1(root *TreeNode) int {
	ret := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		count := queue.Len()
		for i := 0; i < count; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				ret = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return ret
}
