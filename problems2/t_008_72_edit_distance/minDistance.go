package t_008_72_edit_distance

// 解法 1：二维动态规划 —— dp[i][j] 表示 word1 前 i 个字符转成 word2 前 j 个字符的最少操作数
// TODO: 实现
func minDistance(word1 string, word2 string) int {

	word1 = " " + word1
	word2 = " " + word2

	word1len := len(word1)
	word2len := len(word2)

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

	for i := 1; i < word1len; i++ {
		for j := 1; j < word2len; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[word1len-1][word2len-1]

}

// 解法 2：滚动数组 —— 逐行递推，只保留上一行与当前行，把二维表压成一维
// TODO: 实现
// func minDistance1(word1 string, word2 string) int {
// 	return 0
// }
