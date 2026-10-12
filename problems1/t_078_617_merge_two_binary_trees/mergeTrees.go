package t_078_617_merge_two_binary_trees

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

// root1     = []int{1, 3, 2, 5}
// root2     = []int{2, 1, 3, 0, 4, 0, 7}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var traversal func(node1, node2 *TreeNode) *TreeNode
	traversal = func(node1, node2 *TreeNode) *TreeNode {
		if node1 == nil {
			return node2
		}
		if node2 == nil {
			return node1
		}
		node1.Val += node2.Val
		node1.Left = traversal(node1.Left, node2.Left)
		node1.Right = traversal(node1.Right, node2.Right)
		return node1
	}
	return traversal(root1, root2)
}
