package t_013_337_house_robber_iii

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

// 解法 1：树形动态规划 —— 后序遍历，每个节点向上返回「偷当前节点」与「不偷当前节点」两个最大值
// TODO: 实现
func rob(root *TreeNode) int {

	var dfs func(node *TreeNode) (int, int)

	dfs = func(node *TreeNode) (int, int) {

		if node == nil {
			return 0, 0
		}

		leftNot, leftRob := dfs(node.Left)
		rightNot, rightRob := dfs(node.Right)

		notRob := max(leftNot, leftRob) + max(rightNot, rightRob)
		robCurrent := node.Val + leftNot + rightNot

		return notRob, robCurrent

	}

	return max(dfs(root))

}

// 解法 2：带记忆化的递归 —— 用 map 缓存每个节点为根的子树最优值，避免重复访问
// TODO: 实现
// func rob1(root *TreeNode) int {
// 	return 0
// }
