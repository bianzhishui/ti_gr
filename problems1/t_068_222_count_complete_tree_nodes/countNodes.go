package t_068_222_count_complete_tree_nodes

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

func countNodes(root *TreeNode) int {
	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth, rightDepth := 0, 0
		leftNode, rightNode := node.Left, node.Right
		for leftNode != nil {
			leftDepth++
			leftNode = leftNode.Left
		}
		for rightNode != nil {
			rightDepth++
			rightNode = rightNode.Right
		}
		// 遍历到一个满二叉树时就能统计出所有节点的个数
		if leftDepth == rightDepth {
			return (2 << leftDepth) - 1
		}
		return count(node.Left) + count(node.Right) + 1
	}
	return count(root)
}
