package t_060_111_minimum_depth_of_binary_tree

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

func minDepth(root *TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				goto RETURN
			}
		}
		count = len(queue)
		ret++
	}
	return ret
RETURN:
	ret++
	return ret
}

func minDepth1(root *TreeNode) int {
	var (
		traversal func(node *TreeNode) int
		min       func(x, y int) int
	)
	min = func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 如果左子树不为空但是右子树为空 则需要继续遍历左子树
		if node.Left != nil && node.Right == nil {
			return traversal(node.Left) + 1
		} else if node.Left == nil && node.Right != nil {
			return traversal(node.Right) + 1
		}
		// 左右子树都为空时返回1
		return min(traversal(node.Left), traversal(node.Right)) + 1
	}
	return traversal(root)
}
