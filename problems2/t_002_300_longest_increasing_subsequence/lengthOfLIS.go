package t_002_300_longest_increasing_subsequence

// 解法 1：动态规划 O(n^2)
// 定义 dp[i] 为「以 nums[i] 结尾」的最长严格递增子序列长度：
// dp[i] = max(dp[j]) + 1，其中 j < i 且 nums[j] < nums[i]；无满足条件的 j 时 dp[i] = 1。
// 注意是「严格」递增，因此内层判断用 nums[j] < nums[i]（而非 <=），相等元素不能接在后面。
// 最终答案为 dp 数组的最大值；数组长度至少为 1，单元素结果为 1。
// TODO: 实现
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

// 解法 2：贪心 + 二分查找 O(n log n)
// 维护 tails 数组，tails[k] 表示「长度为 k+1 的所有严格递增子序列中，结尾元素的最小值」。
// tails 天然严格递增（否则更短的序列不可能有更大的最小结尾），因此可以二分。
// 遍历每个元素 v，用二分找到 tails 中第一个 >= v 的位置并替换成 v：
// 若能追加到末尾则说明找到了更长的递增子序列，否则用更小的结尾值改进既有长度。
// 注意必须用 lower_bound（第一个 >= v）以保持严格递增；若用 upper_bound 则相等的元素会错误地延长序列。
// tails 的长度即为答案。
// TODO: 实现
// func lengthOfLIS1(nums []int) int {
// 	return 0
// }
