package t_079_700_search_in_a_binary_search_tree

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

// root      = []int{4, 2, 7, 1, 3}
// val       = 2
func searchBST(root *TreeNode, val int) *TreeNode {
	var traversal func(node *TreeNode) *TreeNode
	traversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val > val {
			return traversal(node.Left)
		} else if node.Val < val {
			return traversal(node.Right)
		} else {
			return node
		}
	}
	return traversal(root)
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	var traversal func(node *TreeNode) *TreeNode
	traversal = func(node *TreeNode) *TreeNode {
		if node == nil || node.Val == val {
			return node
		}
		if node.Val > val {
			return traversal(node.Left)
		}
		return traversal(node.Right)
	}
	return traversal(root)
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}
