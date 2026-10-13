package t_005_1143_longest_common_subsequence

// 解法 1：二维动态规划 —— dp[i][j] 表示 text1 前 i 个字符与 text2 前 j 个字符的 LCS 长度
// TODO: 实现
func longestCommonSubsequence(text1 string, text2 string) int {

	text1 = " " + text1
	text2 = " " + text2

	text1len := len(text1)
	text2len := len(text2)

	dp := make([][]int, text1len)
	for i := range dp {
		dp[i] = make([]int, text2len)
	}

	maxLen := 0

	for i := 1; i < text1len; i++ {
		for j := 1; j < text2len; j++ {

			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}

		}
	}

	return maxLen
}

// 解法 2：滚动数组优化空间 —— 先令 text2 为较短串，再把空间压缩到 O(len(text2))
// TODO: 实现
// func longestCommonSubsequence1(text1 string, text2 string) int {
// 	return 0
// }

// 动态规划 一维数组模拟二维数组
func longestCommonSubsequence2(text1 string, text2 string) int {

	text1 = " " + text1
	text2 = " " + text2

	rows := len(text1) // m+1
	cols := len(text2) // n+1

	dp := make([]int, rows*cols)

	// 访问 dp[i][j] 的辅助函数
	at := func(i, j int) int {
		return dp[i*(cols)+j]
	}
	set := func(i, j, v int) {
		dp[i*(cols)+j] = v
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if text1[i] == text2[j] {
				set(i, j, at(i-1, j-1)+1)
			} else {
				set(i, j, max(at(i-1, j), at(i, j-1)))
			}
		}
	}

	return at(rows-1, cols-1) // 等价于 dp[m][n]
}
