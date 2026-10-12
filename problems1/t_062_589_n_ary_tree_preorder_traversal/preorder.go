package t_062_589_n_ary_tree_preorder_traversal

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

func preorder(root *Node) []int {
	ret := []int{}
	var traversal func(node *Node)
	traversal = func(node *Node) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		for _, val := range node.Children {
			traversal(val)
		}
	}
	traversal(root)
	return ret
}
