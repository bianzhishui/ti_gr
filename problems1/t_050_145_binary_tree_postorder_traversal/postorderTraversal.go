package t_050_145_binary_tree_postorder_traversal

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

func postorderTraversal(root *TreeNode) []int {
	ret := []int{}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		ret = append(ret, root.Val)
	}
	traversal(root)
	return ret
}

func postorderTraversal1(root *TreeNode) []int {
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
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-i-1] = ret[l-i-1], ret[i]
	}
	return ret
}

func postorderTraversal2(root *TreeNode) []int {
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
			ret = append(ret, node.Val)
		}
	}
	return ret
}
