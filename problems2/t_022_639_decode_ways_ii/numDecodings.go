package t_022_639_decode_ways_ii

// 解法 1：线性动态规划
// dp[i] 表示前缀 s[:i] 的解码方案数，从左到右滚动转移：
// 一位解码时按 s[i] 是 '*'、'0' 还是普通数字分别贡献 9、0、1 倍；两位解码时按 s[i-1]、s[i] 的组合分类，
// '*' 与数字的搭配分别对应 15、2、1、9、6 等倍数，所有累加都对 1e9+7 取模。
// TODO: 实现
func numDecodings(s string) int {
	return 0
}

// 解法 2：记忆化搜索
// 自顶向下递归求「剩余子串 s[i:] 的解码方案数」，把 s[i] 单独成码与 s[i:i+2] 组合成码两条分支的结果累加，
// s[i] 为 '0' 且无法与前一位组合时该分支返回 0，用 memo 缓存已算过的下标避免重复展开。
// TODO: 实现
func numDecodings1(s string) int {
	return 0
}
