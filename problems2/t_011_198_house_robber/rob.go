package t_011_198_house_robber

import "slices"

// 解法 1：记忆化搜索 —— 自顶向下递归，缓存从第 i 间房开始能抢到的最大值
// TODO: 实现
// func rob(nums []int) int {
// 	return 0
// }

// 解法 2：动态规划 —— 自底向上递推，逐间房转移并只保留最近两个状态
// TODO: 实现
func rob1(nums []int) int {

	nums = slices.Insert(nums, 0, 0)
	nums = slices.Insert(nums, 0, 0)

	numsLen := len(nums)

	dp := make([]int, numsLen)

	for i := 2; i < numsLen; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[numsLen-1]

}
