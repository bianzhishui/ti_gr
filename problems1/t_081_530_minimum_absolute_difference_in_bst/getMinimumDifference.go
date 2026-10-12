package t_081_530_minimum_absolute_difference_in_bst

import (
	"container/list"
	"math"
)

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

func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt64
	var pre *TreeNode
	var traversal func(cur *TreeNode)
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Left)
		if pre != nil {
			result = min(result, cur.Val-pre.Val)
		}
		pre = cur
		traversal(cur.Right)
	}
	traversal(root)
	return result
}

// 迭代
func getMinimumDifference1(root *TreeNode) int {
	// 返回值
	result := math.MaxInt64
	// 记录上一个节点位置
	var pre *TreeNode
	// 返回最小值
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	// 游标
	cur := root
	// 栈
	stack := list.New()
	// 遍历
	for cur != nil || stack.Len() != 0 {
		// 当前节点不为nil 压栈 继续遍历左节点
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			// cur 为nil 说明左节点遍历完 该处理中间结点
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil {
				result = min(result, cur.Val-pre.Val)
			}
			pre = cur
			// 遍历右节点
			cur = cur.Right
		}
	}
	return result
}
