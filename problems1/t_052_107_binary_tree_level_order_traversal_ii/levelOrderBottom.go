package t_052_107_binary_tree_level_order_traversal_ii

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

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		nums := []int{}
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			nums = append(nums, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count = len(queue)
		ret = append(ret, nums)
	}
	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-1-i] = ret[l-1-i], ret[i]
	}
	return ret
}
