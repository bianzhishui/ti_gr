package t_077_654_maximum_binary_tree

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

// nums = []int{3, 2, 1, 6, 0, 5}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var getMaxNum func(nums []int) int
	getMaxNum = func(nums []int) int {
		l := len(nums)
		if l == 0 {
			return -1
		}
		index := 0
		for i := 1; i < l; i++ {
			if nums[i] > nums[index] {
				index = i
			}
		}
		return index
	}
	var getIndexByNums func(nums []int) *TreeNode
	getIndexByNums = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		index := getMaxNum(nums)
		return &TreeNode{
			Val:   nums[index],
			Left:  getIndexByNums(nums[:index]),
			Right: getIndexByNums(nums[index+1:]),
		}
	}
	return getIndexByNums(nums)
}
