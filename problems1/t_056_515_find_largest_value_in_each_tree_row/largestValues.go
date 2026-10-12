package t_056_515_find_largest_value_in_each_tree_row

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

func largestValues(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 1
	var max int
	for len(queue) != 0 {
		max = queue[0].Val
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count = len(queue)
		ret = append(ret, max)
	}
	return ret
}
