package t_075_106_construct_binary_tree_from_inorder_and_postorder_traversal

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

// inorder   = []int{9, 3, 15, 20, 7}
// postorder = []int{9, 15, 7, 20, 3}
func buildTree(inorder []int, postorder []int) *TreeNode {
	var searchSlice func(slice []int, num int) int
	searchSlice = func(slice []int, num int) int {
		for index, val := range slice {
			if val == num {
				return index
			}
		}
		return -1
	}
	var build func(inorder []int, postorder []int) *TreeNode
	build = func(inorder, postorder []int) *TreeNode {
		if len(inorder) == 0 || len(postorder) == 0 {
			return nil
		}
		root := postorder[len(postorder)-1]
		index := searchSlice(inorder, root)
		return &TreeNode{
			Val:   root,
			Left:  build(inorder[:index], postorder[:index]),
			Right: build(inorder[index+1:], postorder[index:len(postorder)-1]),
		}
	}
	return build(inorder, postorder)
}
