package t_010_392_is_subsequence

// 解法 1：双指针贪心 —— 用一个指针从左到右扫 t，按顺序匹配 s 的每个字符，能匹配完即为子序列
// TODO: 实现
// func isSubsequence(s string, t string) bool {
// 	return false
// }

// 解法 2：动态规划 —— 预处理每个位置之后各字符首次出现的位置，之后对每个 s 只需顺着跳转，适合大批量查询
// TODO: 实现
func isSubsequence1(s string, t string) bool {

	s = " " + s
	t = " " + t

	slen := len(s)
	tlen := len(t)

	dp := make([][]int, slen)
	for i := range dp {
		dp[i] = make([]int, tlen)
		if i == 0 {
			for j := 0; j < tlen; j++ {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < slen; i++ {
		for j := 1; j < tlen; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[slen-1][tlen-1] > 0

}
