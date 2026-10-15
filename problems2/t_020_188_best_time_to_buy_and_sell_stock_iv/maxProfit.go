package t_020_188_best_time_to_buy_and_sell_stock_iv

// 解法 1：三维动态规划
// 状态 dp[天数][交易次数][是否持有]，逐天按「不动 / 买入 / 卖出」转移。
// TODO: 实现
func maxProfit(k int, prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	const inf = -1 << 30

	// dp[i][k][j]：第 i 天、完成 k 笔交易、j=0/1 不持有/持有 的最大利润
	dp := make([][][2]int, n)
	for i := range dp {

		dp[i] = make([][2]int, k+1)

		if i == 0 {
			for kk := 0; kk <= k; kk++ {
				dp[i][kk][0] = 0
				dp[i][kk][1] = -prices[0]
			}
		}
	}
	dp[0][0][1] = inf // 无交易额度却持股，非法

	for i := 1; i < n; i++ {
		for kk := 0; kk <= k; kk++ {
			// 不持股：不动 or 卖出
			dp[i][kk][0] = max(dp[i-1][kk][0], dp[i-1][kk][1]+prices[i])
			// 持股：不动 or 买入（消耗一笔额度）
			if kk > 0 {
				dp[i][kk][1] = max(dp[i-1][kk][1], dp[i-1][kk-1][0]-prices[i])
			} else {
				dp[i][kk][1] = dp[i-1][kk][1] // k=0 不能买入
			}
		}
	}
	return dp[n-1][k][0]

}

// 解法 2：滚动数组压缩天数维度
// 只保留上一层的 dp[交易次数][是否持有]，用两个二维数组交替更新。
// TODO: 实现
// func maxProfit1(k int, prices []int) int {
// 	return 0
// }
