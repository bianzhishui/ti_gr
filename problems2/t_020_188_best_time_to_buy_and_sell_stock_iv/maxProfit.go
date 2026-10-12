package t_020_188_best_time_to_buy_and_sell_stock_iv

// 解法 1：三维动态规划
// 状态 dp[天数][交易次数][是否持有]，逐天按「不动 / 买入 / 卖出」转移。
// TODO: 实现
func maxProfit(k int, prices []int) int {
	return 0
}

// 解法 2：滚动数组压缩天数维度
// 只保留上一层的 dp[交易次数][是否持有]，用两个二维数组交替更新。
// TODO: 实现
func maxProfit1(k int, prices []int) int {
	return 0
}
