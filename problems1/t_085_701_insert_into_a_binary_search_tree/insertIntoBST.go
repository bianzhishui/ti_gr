package t_085_701_insert_into_a_binary_search_tree

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var traversal func(cur *TreeNode) *TreeNode
	traversal = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			cur = &TreeNode{
				Val: val,
			}
			return cur
		}
		if val > cur.Val {
			cur.Right = traversal(cur.Right)
		} else if val < cur.Val {
			cur.Left = traversal(cur.Left)
		}
		return cur
	}
	return traversal(root)
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}
	if root == nil {
		root = newNode
		return root
	}
	var pre *TreeNode
	cur := root
	for cur != nil {
		pre = cur
		if cur.Val > val {
			cur = cur.Left
		} else if cur.Val < val {
			cur = cur.Right
		}
	}
	if pre.Val > val {
		pre.Left = newNode
	} else if pre.Val < val {
		pre.Right = newNode
	}
	return root
}
