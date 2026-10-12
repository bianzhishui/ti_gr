package t_074_113_path_sum_ii

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

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := [][]int{}
	var traversal func(node *TreeNode, sum int, path []int)
	traversal = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == node.Val {
			// copy结果集 将copy出来的切片添加到ret中 原切片底层数组的值会在程序运行中发生改变
			pathCopy := make([]int, 0, len(path))
			for _, val := range path {
				pathCopy = append(pathCopy, val)
			}
			ret = append(ret, pathCopy)
			return
		}
		sum -= node.Val
		traversal(node.Left, sum, path)
		traversal(node.Right, sum, path)
	}
	traversal(root, targetSum, []int{})
	return ret
}
