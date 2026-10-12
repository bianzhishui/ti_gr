package t_088_108_convert_sorted_array_to_binary_search_tree

import "container/list"

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

// 递归
func sortedArrayToBST(nums []int) *TreeNode {
	var createNode func(nums []int) *TreeNode
	createNode = func(nums []int) *TreeNode {
		l := len(nums)
		if l == 0 {
			return nil
		}
		mid := l / 2
		node := &TreeNode{
			Val: nums[mid],
		}
		node.Left = createNode(nums[:mid])
		node.Right = createNode(nums[mid+1:])
		return node
	}
	return createNode(nums)
}

// 迭代
func sortedArrayToBST1(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	// left和right为数组边界[left,right]
	type nodeInfo struct {
		node  *TreeNode
		left  int
		right int
	}
	// 队列
	nodeInfoQueue := list.New()
	// 返回的根节点
	root := &TreeNode{}
	nodeInfoQueue.PushBack(&nodeInfo{
		node:  root,
		left:  0,
		right: l - 1,
	})
	// 队列不为空循环
	for nodeInfoQueue.Len() != 0 {
		// 弹出队列首元素
		nodeinfo := nodeInfoQueue.Remove(nodeInfoQueue.Front()).(*nodeInfo)
		// 找到数组中间值
		mid := nodeinfo.left + (nodeinfo.right-nodeinfo.left)/2
		// 赋值
		nodeinfo.node.Val = nums[mid]
		// 如果在nums切片中mid左边有元素
		if nodeinfo.left < mid {
			node := &TreeNode{}
			nodeinfo.node.Left = node
			nodeInfoQueue.PushBack(&nodeInfo{
				node:  node,
				left:  nodeinfo.left,
				right: mid - 1,
			})
		}
		// 如果在nums切片中mid右边有元素
		if nodeinfo.right > mid {
			node := &TreeNode{}
			nodeinfo.node.Right = node
			nodeInfoQueue.PushBack(&nodeInfo{
				node:  node,
				left:  mid + 1,
				right: nodeinfo.right,
			})
		}
	}
	return root
}
