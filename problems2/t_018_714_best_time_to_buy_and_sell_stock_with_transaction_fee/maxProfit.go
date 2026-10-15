package t_018_714_best_time_to_buy_and_sell_stock_with_transaction_fee

import "slices"

// 解法 1：状态机动态规划
// 维护「持有股票」与「不持有股票」两个状态，卖出时扣手续费。
// TODO: 实现
func maxProfit(prices []int, fee int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	dp1 := make([]int, n) // 当天结束 持仓
	dp2 := make([]int, n) // 当天结束 空仓

	dp1[0] = -prices[0]

	for i := 1; i < n; i++ {

		dp1[i] = max(dp1[i-1], dp2[i-1]-prices[i])
		dp2[i] = max(dp2[i-1], dp1[i-1]+prices[i]-fee)

	}

	return slices.Max(dp2)

}

// 解法 2：滚动变量压缩状态
// 用两个滚动变量滚动更新两个状态，把空间降到 O(1)。
// TODO: 实现
// func maxProfit1(prices []int, fee int) int {
// 	return 0
// }
