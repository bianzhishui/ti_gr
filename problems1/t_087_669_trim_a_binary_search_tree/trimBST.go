package t_087_669_trim_a_binary_search_tree

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var traversal func(cur *TreeNode) *TreeNode
	traversal = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		if cur.Val < low {
			return traversal(cur.Right)
		}
		if cur.Val > high {
			return traversal(cur.Left)
		}
		cur.Left = traversal(cur.Left)
		cur.Right = traversal(cur.Right)
		return cur
	}
	return traversal(root)
}

// 迭代
func trimBST1(root *TreeNode, low int, high int) *TreeNode {
	// 找到根节点的位置
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else if root.Val > high {
			root = root.Left
		}
	}
	// 找到的根节点为nil返回
	if root == nil {
		return nil
	}
	// 修剪左子树
	for cur := root; cur.Left != nil; {
		// 处理左节点
		if cur.Left.Val < low {
			cur.Left = cur.Left.Right
		} else {
			cur = cur.Left
		}
	}
	// 修剪右子树
	for cur := root; cur.Right != nil; {
		// 处理右节点
		if cur.Right.Val > high {
			cur.Right = cur.Right.Left
		} else {
			cur = cur.Right
		}
	}
	return root
}
