package t_055_429_n_ary_tree_level_order_traversal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*Node{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		nums := []int{}
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			nums = append(nums, node.Val)
			if node.Children != nil {
				queue = append(queue, node.Children...)
			}
		}
		count = len(queue)
		ret = append(ret, nums)
	}
	return ret
}
