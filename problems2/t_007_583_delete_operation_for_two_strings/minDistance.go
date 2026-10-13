package t_007_583_delete_operation_for_two_strings

// 解法 1：先求最长公共子序列，答案 = m + n - 2 * LCS
// TODO: 实现
// func minDistance(word1 string, word2 string) int {
// 	return 0
// }

// 解法 2：直接二维动态规划：dp[i][j] 表示使前 i、前 j 个字符相同所需的最少删除次数
// TODO: 实现
func minDistance1(word1 string, word2 string) int {

	word1 = " " + word1
	word2 = " " + word2

	word1len := len(word1)
	word2len := len(word2)

	// 定义dp并初始化边界
	dp := make([][]int, word1len)
	for i := range dp {
		dp[i] = make([]int, word2len)
		if i == 0 {
			for j := 0; j < word2len; j++ {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}

	result := 0

	for i := 1; i < word1len; i++ {
		for j := 1; j < word2len; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
			}
			result = dp[i][j]
		}
	}

	return result
}
