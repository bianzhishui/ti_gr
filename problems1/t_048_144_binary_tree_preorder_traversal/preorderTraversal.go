package t_048_144_binary_tree_preorder_traversal

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

//	  0
//	1    2
//
// 3   4 5  6

// 0 1 3 4 2 5 6
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return ret
}

func preorderTraversal1(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		node := stack[l-1]
		if node != nil {
			stack = stack[:l-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			stack = stack[:l-1]
			node = stack[l-2]
			stack = stack[:l-2]
			ret = append(ret, node.Val)
		}
	}
	return ret
}
