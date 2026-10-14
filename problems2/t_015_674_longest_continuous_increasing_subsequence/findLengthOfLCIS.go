package t_015_674_longest_continuous_increasing_subsequence

import "slices"

// 解法 1：一次遍历
// 顺序扫描数组，用一个变量记录「以当前位置结尾」的连续递增段长度：
// 若 nums[i] > nums[i-1] 则计数加一，否则把计数重置为 1，全程维护最大值。
// 元素相等不算递增，需要重置；数组长度至少为 1，单元素答案为 1。
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur, result := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			cur = 1
		}

		if cur > result {
			result = cur
		}
	}

	return result
}

// 解法 2：动态规划
// dp[i] 表示以 nums[i] 结尾的最长连续递增序列长度：
// 若 nums[i] > nums[i-1] 则 dp[i] = dp[i-1] + 1，否则 dp[i] = 1；答案为 dp 中的最大值。
// 与「最长递增子序列」不同，这里要求下标连续，因此只需与前一个元素比较，无需回看更早的位置。
// TODO: 实现
func findLengthOfLCIS1(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	dp := make([]int, numsLen)
	dp[0] = 1

	for i := 1; i < numsLen; i++ {

		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}

	}

	return slices.Max(dp)

}
