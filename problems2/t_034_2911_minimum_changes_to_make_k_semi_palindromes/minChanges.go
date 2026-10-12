package t_034_2911_minimum_changes_to_make_k_semi_palindromes

// 解法 1：预处理每个子串变成半回文串的最少修改次数 + 划分型动态规划（切成 k 段）
// 先对每个子串枚举其全部合法因数，求出把它变成半回文串的最少修改次数；
// 再按「前缀切成 j 段」自底向上做划分型 DP，每段长度至少为 2。
// TODO: 实现
func minChanges(s string, k int) int {
	return 0
}

// 解法 2：记忆化搜索：自顶向下按 (起点, 剩余段数) 递归
// 复用同一份子串修改次数预处理表，从起点枚举本段结束位置，用 (起点, 剩余段数) 记忆化。
// TODO: 实现
func minChanges1(s string, k int) int {
	return 0
}
