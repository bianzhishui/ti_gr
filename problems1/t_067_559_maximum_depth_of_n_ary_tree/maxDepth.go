package t_067_559_maximum_depth_of_n_ary_tree

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

func maxDepth(root *Node) int {
	var traversal func(node *Node) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	traversal = func(node *Node) int {
		if node == nil {
			return 0
		}
		num := 0
		for _, val := range node.Children {
			num = max(traversal(val), num)
		}
		return num + 1
	}
	return traversal(root)
}
