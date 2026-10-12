package t_084_235_lowest_common_ancestor_of_a_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			break
		}
	}
	return cur
}

// 递归
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var traversal func(cur *TreeNode) *TreeNode
	traversal = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		if p.Val > cur.Val && q.Val > cur.Val {
			return traversal(cur.Right)
		} else if p.Val < cur.Val && q.Val < cur.Val {
			return traversal(cur.Left)
		} else {
			return cur
		}
	}
	return traversal(root)
}
