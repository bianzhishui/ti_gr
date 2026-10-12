package t_054_637_average_of_levels_in_binary_tree

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

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		var sum float64
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, sum/float64(count))
		count = len(queue)
	}
	return ret
}
