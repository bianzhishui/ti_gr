package t_083_236_lowest_common_ancestor_of_a_binary_tree

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

// 递归遍历两次统计出数组 从数组里取最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 获取一个节点的所有祖先包括它自己
	var getAllAncestor func(cur, node *TreeNode) []*TreeNode
	getAllAncestor = func(cur, node *TreeNode) []*TreeNode {
		var (
			// 返回切片
			allAncestor []*TreeNode
			// 前序遍历
			traversal func(cur *TreeNode) bool
		)
		traversal = func(cur *TreeNode) bool {
			// 节点为nil说明该路径没有找到 返回false
			if cur == nil {
				return false
			}
			// 中
			// 将遍历到的该节点加入返回切片中
			allAncestor = append(allAncestor, cur)
			// 若该节点为目标节点 返回true
			if cur == node {
				return true
			}
			// 左
			if traversal(cur.Left) {
				return true
			}
			// 右
			if traversal(cur.Right) {
				return true
			}
			// 回溯
			allAncestor = allAncestor[:len(allAncestor)-1]
			return false
		}
		traversal(cur)
		return allAncestor
	}
	// 获取pq节点的所有祖先分别组成数组
	pAllAncestor := getAllAncestor(root, p)
	qAllAncestor := getAllAncestor(root, q)
	// 从数组中找到pq节点的最近公共祖先
	var cur *TreeNode
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	l := min(len(pAllAncestor), len(qAllAncestor))
	for i := 0; i < l; i++ {
		if pAllAncestor[i] != qAllAncestor[i] {
			break
		}
		cur = pAllAncestor[i]
	}
	return cur
}

// 递归后序遍历
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var traversal func(cur, p, q *TreeNode) *TreeNode
	traversal = func(cur, p, q *TreeNode) *TreeNode {
		if cur == nil || cur == p || cur == q {
			return cur
		}
		left := traversal(cur.Left, p, q)
		right := traversal(cur.Right, p, q)
		if left != nil && right == nil {
			return left
		} else if left == nil && right != nil {
			return right
		} else if left != nil && right != nil {
			return cur
		} else {
			return nil
		}
	}
	return traversal(root, p, q)
}
