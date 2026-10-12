package t_076_105_construct_binary_tree_from_preorder_and_inorder_traversal

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
// postorder = []int{3, 9, 20, 15, 7}
func buildTree(preorder []int, inorder []int) *TreeNode {
	var searchIndexBySlice func(slicer []int, num int) int
	searchIndexBySlice = func(slicer []int, num int) int {
		for index, val := range slicer {
			if val == num {
				return index
			}
		}
		return -1
	}
	var build func(preorder, inorder []int) *TreeNode
	build = func(preorder, inorder []int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {
			return nil
		}
		root := preorder[0]
		index := searchIndexBySlice(inorder, root)
		return &TreeNode{
			Val:   root,
			Left:  build(preorder[1:index+1], inorder[:index]),
			Right: build(preorder[index+1:], inorder[index+1:]),
		}
	}
	return build(preorder, inorder)
}
