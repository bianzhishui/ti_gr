package t_065_100_same_tree

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var compare func(p, q *TreeNode) bool
	compare = func(p, q *TreeNode) bool {
		if p == nil && q != nil {
			return false
		} else if p != nil && q == nil {
			return false
		} else if p == nil && q == nil {
			return true
		} else if p.Val != q.Val {
			return false
		} else {
			return compare(p.Left, q.Left) && compare(p.Right, q.Right)
		}
	}
	return compare(p, q)
}
