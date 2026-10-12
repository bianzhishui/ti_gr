package t_033_1043_partition_array_for_maximum_sum

// 解法 1：线性动态规划
// dp[i] 表示前缀 arr[:i] 分隔后能得到的元素最大和，枚举最后一段的长度 j（1 <= j <= k），
// 用该段的最大值替换整段后累加到 dp[i-j] 上，取所有 j 中的最大值。
// TODO: 实现
func maxSumAfterPartitioning(arr []int, k int) int {
	return 0
}

// 解法 2：记忆化搜索
// 自顶向下按「已处理长度 i」递归，枚举末尾那一段的长度 j（1 <= j <= k），
// 累加「该段最大值 × 段长」并取最优；用 memo 缓存已经算过的长度，避免重复展开。
// TODO: 实现
func maxSumAfterPartitioning1(arr []int, k int) int {
	return 0
}
