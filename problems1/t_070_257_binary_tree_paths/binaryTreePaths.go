package t_070_257_binary_tree_paths

import "strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	if root == nil {
		return ret
	}
	var traversal func(node *TreeNode, s string)
	traversal = func(node *TreeNode, s string) {
		// 找到末尾节点
		if node.Left == nil && node.Right == nil {
			s += strconv.Itoa(node.Val)
			ret = append(ret, s)
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			traversal(node.Left, s)
		}
		if node.Right != nil {
			traversal(node.Right, s)
		}
	}
	traversal(root, "")
	return ret
}

func binaryTreePaths1(root *TreeNode) []string {
	// 栈
	stack := []*TreeNode{}
	paths := make([]string, 0)
	// 返回值
	res := make([]string, 0)
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}
	for len(stack) > 0 {
		l := len(stack)
		// 栈顶元素
		node := stack[l-1]
		path := paths[l-1]
		// 弹出栈顶元素
		stack = stack[:l-1]
		paths = paths[:l-1]
		// 末尾节点
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}
		// 右节点入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		// 左节点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}
