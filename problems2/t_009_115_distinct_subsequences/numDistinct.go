package t_009_115_distinct_subsequences

// 解法 1：二维动态规划计数 —— dp[i][j] 表示 s 前 i 个字符中 t 前 j 个字符出现的个数
// TODO: 实现
func numDistinct(s string, t string) int {

	s = " " + s
	t = " " + t

	slen := len(s)
	tlen := len(t)

	dp := make([][]int, slen)
	for i := range dp {
		dp[i] = make([]int, tlen)
		dp[i][0] = 1
	}

	for i := 1; i < slen; i++ {
		for j := 1; j < tlen; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[slen-1][tlen-1]

}

// 解法 2：滚动数组把空间优化到一维（内层需倒序，避免复用本轮已更新的值）
// TODO: 实现
// func numDistinct1(s string, t string) int {
// 	return 0
// }
