package t_058_117_populating_next_right_pointers_in_each_node_ii

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		var cur *Node
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i != 0 {
				cur.Next = node
			}
			cur = node
		}
		count = len(queue)
	}
	return root
}
