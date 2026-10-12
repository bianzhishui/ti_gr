package t_063_590_n_ary_tree_postorder_traversal

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

func postorder(root *Node) []int {
	ret := []int{}
	var traversal func(node *Node)
	traversal = func(node *Node) {
		if node == nil {
			return
		}
		for _, val := range node.Children {
			traversal(val)
		}
		ret = append(ret, node.Val)
	}
	traversal(root)
	return ret
}
