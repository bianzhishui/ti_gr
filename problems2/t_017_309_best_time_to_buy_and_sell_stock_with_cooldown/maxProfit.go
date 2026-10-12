package t_017_309_best_time_to_buy_and_sell_stock_with_cooldown

// 解法 1：状态机动态规划
// 按天递推，每天区分「持有股票」「不持有且不在冷冻期」「当天卖出进入冷冻期」三种状态，逐日转移取最大值。
// TODO: 实现
func maxProfit(prices []int) int {
	return 0
}

// 解法 2：滚动变量压缩状态
// 状态定义与解法 1 相同，但只保留前一天的三个状态量，用常量个变量滚动更新。
// TODO: 实现
func maxProfit1(prices []int) int {
	return 0
}
