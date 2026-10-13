package t_003_673_number_of_longest_increasing_subsequence

// 解法 1：动态规划 —— 同时维护以 i 结尾的最长长度与达到该长度的方案数
// TODO: 实现
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)    // dp[i]: 以 nums[i] 结尾的最长递增子序列长度
	count := make([]int, n) // count[i]: 以 nums[i] 结尾的最长递增子序列个数
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}

	maxLen := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					// 找到了更长的子序列，重置计数
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					// 找到相同长度的新子序列，累加计数
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	// 统计所有以最长长度结尾的子序列个数
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxLen {
			result += count[i]
		}
	}
	return result
}

// 解法 2：树状数组按值域维护长度与计数
// TODO: 实现
// func findNumberOfLIS1(nums []int) int {
// 	return 0
// }
