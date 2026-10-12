package t_069_110_balanced_binary_tree

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

func isBalanced(root *TreeNode) bool {
	var getHeight func(node *TreeNode) int
	var abs func(x int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	abs = func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	// 左右中 后序遍历获取高度
	getHeight = func(node *TreeNode) int {
		// 递归结束条件
		if node == nil {
			return 0
		}
		// 获取左叶子节点高度
		leftHeight := getHeight(node.Left)
		if leftHeight == -1 {
			return -1
		}
		// 获取右叶子节点高度
		rightHeight := getHeight(node.Right)
		if rightHeight == -1 {
			return -1
		}
		// 如果左右节点高度差大于1 说明不是平衡树
		if abs(leftHeight-rightHeight) > 1 {
			return -1
			// 否则返回高度
		} else {
			return 1 + max(leftHeight, rightHeight)
		}
	}
	// 返回-1的时候说明不是平衡树
	if getHeight(root) == -1 {
		return false
	} else {
		return true
	}
}
