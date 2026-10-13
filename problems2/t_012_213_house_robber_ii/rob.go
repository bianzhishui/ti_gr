package t_012_213_house_robber_ii

import "slices"

// 解法 1：环形拆成两段线性问题 —— 首尾不能同时偷，分别考虑不含首元素与不含尾元素，取较大值
// TODO: 实现
func rob(nums []int) int {

	robFunc := func(nums []int) int {

		nums = slices.Insert(nums, 0, 0, 0)

		numsLen := len(nums)

		dp := make([]int, numsLen)

		for i := 2; i < numsLen; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}

		return dp[numsLen-1]

	}

	numsLen := len(nums)

	if numsLen > 1 {
		return max(robFunc(slices.Delete(slices.Clone(nums), 0, 1)), robFunc(slices.Delete(slices.Clone(nums), numsLen-1, numsLen)))
	} else {
		return robFunc(nums)
	}

}

// 解法 2：记忆化搜索 —— 对上述两段分别做带缓存的递归
// TODO: 实现
// func rob1(nums []int) int {
// 	return 0
// }
