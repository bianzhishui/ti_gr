package t_017_309_best_time_to_buy_and_sell_stock_with_cooldown

/*

        买入                卖出
  空仓 ────────→ 持有 ────────→ 刚卖出
   ↑                              │
   │         冷冻期结束            │
   └──────────────────────────────┘

*/

// 解法 1：状态机动态规划
// 按天递推，每天区分「持有股票」「不持有且不在冷冻期」「当天卖出进入冷冻期」三种状态，逐日转移取最大值。
// TODO: 实现
func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	// 第 i 天结束时，持有股票的最大收益
	hold := make([]int, n)
	// 第 i 天结束时，不持有股票，且今天卖出了的最大收益（即处于冷冻期）
	sold := make([]int, n)
	// 第 i 天结束时，不持有股票，且今天没卖出（即不在冷冻期，可以随时买入）的最大收益
	rest := make([]int, n)

	hold[0] = -prices[0]

	for i := 1; i < n; i++ {

		hold[i] = max(hold[i-1], rest[i-1]-prices[i]) // 持有 前一天空仓买入 或 未买入也未卖出
		sold[i] = hold[i-1] + prices[i]               // 刚卖出 前一天持有+当天股票价值
		rest[i] = max(rest[i-1], sold[i-1])           // 空仓 前一天也空仓 或 前一天卖出

	}

	return max(sold[n-1], rest[n-1])

}

// 解法 2：滚动变量压缩状态
// 状态定义与解法 1 相同，但只保留前一天的三个状态量，用常量个变量滚动更新。
// TODO: 实现
// func maxProfit1(prices []int) int {
// 	return 0
// }
