package t_061_226_invert_binary_tree

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

func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		invert(node.Left)
		invert(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	invert(root)
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		node := stack[l-1]
		if node != nil {
			stack = stack[:l-1]
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			stack = stack[:l-1]
			node = stack[l-2]
			stack = stack[:l-2]
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}
