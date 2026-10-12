package t_086_450_delete_node_in_a_bst

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
func deleteNode(root *TreeNode, key int) *TreeNode {
	var traversal func(root *TreeNode) *TreeNode
	traversal = func(root *TreeNode) *TreeNode {
		// 节点为空返回nil
		if root == nil {
			return nil
		}
		// 若当前节点值于key相同 说明找到了删除的节点
		if root.Val == key {
			// 删除的节点为叶子节点直接返回nil
			if root.Left == nil && root.Right == nil {
				return nil
				// 删除的节点右节点为nil 则将该节点的左节点返回
			} else if root.Right == nil {
				return root.Left
				// 删除的节点左节点为nil 则将该节点的右节点返回
			} else if root.Left == nil {
				return root.Right
				// 若删除的节点 左右子节点不为空
				// 则 找到删除节点的右节点的左子树最左位置
			} else {
				cur := root.Right
				for cur.Left != nil {
					cur = cur.Left
				}
				// 将删除节点的左子树接到找到的位置
				cur.Left = root.Left
				return root.Right
			}
			// 若当前节点值大于key向左继续遍历
		} else if root.Val > key {
			root.Left = traversal(root.Left)
			// 若当前节点值小于key向右继续遍历
		} else {
			root.Right = traversal(root.Right)
		}
		// 返回处理完的当前节点
		return root
	}
	return traversal(root)
}

// 迭代
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var pre *TreeNode
	cur := root
	// 真正删除节点的逻辑
	var delOneNode = func() *TreeNode {
		// 叶子节点的情况
		if cur.Left == nil && cur.Right == nil {
			return nil
			// 左节点为空的情况
		} else if cur.Left == nil {
			return cur.Right
			// 右节点为空的情况
		} else if cur.Right == nil {
			return cur.Left
			// 左右节点不为空的情况
		} else {
			curTmp := cur.Right
			for curTmp.Left != nil {
				curTmp = curTmp.Left
			}
			curTmp.Left = cur.Left
			return cur.Right
		}
	}
	var modifyPre = func() {
		// 若pre为nil 说明删除的是根节点 所以修改root
		if pre == nil {
			root = delOneNode()
			return
		}
		// 找到需要修改的是左节点还是右节点
		if pre.Left != nil && pre.Left.Val == key {
			pre.Left = delOneNode()
		} else {
			pre.Right = delOneNode()
		}
	}
	// 找到删除的节点
	for cur != nil {
		if cur.Val == key {
			modifyPre()
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return root
}
