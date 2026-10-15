package t_019_123_best_time_to_buy_and_sell_stock_iii

import (
	"math"
)

// 解法 1：状态机动态规划 —— 把「第几次交易 × 是否持有」拆成 5 个状态线性转移
// TODO: 实现
func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	const k = 2
	const inf = -1 << 30

	// dp[i][k][j]：第 i 天、完成 k 笔交易、j=0/1 不持有/持有 的最大利润
	dp := make([][k + 1][2]int, n)
	for kk := 0; kk <= k; kk++ {
		dp[0][kk][0] = 0
		dp[0][kk][1] = -prices[0]
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

// 解法 2：滚动变量压缩状态 —— 用常数个变量滚动更新，把空间降到 O(1)
// TODO: 实现
// func maxProfit1(prices []int) int {
// 	return 0
// }

func maxProfit2(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	const negInf = math.MinInt32 / 2

	dp00 := make([]int, n) // 空仓0
	dp11 := make([]int, n) // 持仓1
	dp10 := make([]int, n) // 空仓1
	dp21 := make([]int, n) // 持仓2
	dp20 := make([]int, n) // 空仓2

	dp11[0] = -prices[0]
	dp10[0] = negInf
	dp21[0] = negInf
	dp20[0] = negInf

	for i := 1; i < n; i++ {

		dp11[i] = max(dp11[i-1], dp00[i-1]-prices[i])
		dp10[i] = max(dp10[i-1], dp11[i-1]+prices[i])
		dp21[i] = max(dp21[i-1], dp10[i-1]-prices[i])
		dp20[i] = max(dp20[i-1], dp21[i-1]+prices[i])

	}

	return max(dp10[n-1], dp20[n-1])

}
