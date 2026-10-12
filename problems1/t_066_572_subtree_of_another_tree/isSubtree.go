package t_066_572_subtree_of_another_tree

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

// 递归检查两个树是否相同
func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.Val != b.Val {
		return false
	} else {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}
	// var check func(a, b *TreeNode) bool
	// 递归检查两个树是否相同
	// check = func(a, b *TreeNode) bool {
	// 	if a == nil && b == nil {
	// 		return true
	// 	} else if a == nil || b == nil {
	// 		return false
	// 	} else if a.Val != b.Val {
	// 		return false
	// 	} else {
	// 		return check(a.Left, b.Left) && check(a.Right, b.Right)
	// 	}
	// }
	// 先判断自身是否与subroot相同
	// 若不同则拿左子树和右子树与subroot比较
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
