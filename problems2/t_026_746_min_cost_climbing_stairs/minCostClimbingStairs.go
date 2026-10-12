package t_026_746_min_cost_climbing_stairs

// 解法 1：动态规划 + 滚动变量 —— 以 i 作为最后一级付费台阶时 dp[i] = cost[i] + min(dp[i-1], dp[i-2])，
// 答案取最后两级的最小值，自底向上递推过程中只保留前两个状态
// TODO: 实现
func minCostClimbingStairs(cost []int) int {
	return 0
}

// 解法 2：记忆化搜索 —— 自顶向下按台阶下标递归，把已经算出的状态缓存起来避免重复计算
// TODO: 实现
func minCostClimbingStairs1(cost []int) int {
	return 0
}
