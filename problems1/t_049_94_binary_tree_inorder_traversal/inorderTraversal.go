package t_049_94_binary_tree_inorder_traversal

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

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		ret = append(ret, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return ret
}

func inorderTraversal1(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) != 0 {
		l := len(stack)
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[l-1]
			stack = stack[:l-1]
			ret = append(ret, node.Val)
			cur = node.Right
		}
	}
	return ret
}

// 迭代
func inorderTraversal2(root *TreeNode) []int {
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
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			stack = stack[:l-1]
			node = stack[l-2]
			stack = stack[:l-2]
			ret = append(ret, node.Val)
		}
	}
	return ret
}
