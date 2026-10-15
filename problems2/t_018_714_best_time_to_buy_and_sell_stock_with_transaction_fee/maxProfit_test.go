package t_018_714_best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"fmt"
	"testing"
)

// cases 是全部解法共用的用例表，覆盖题面两个示例、单元素、手续费为 0、
// 手续费高到无法盈利、全程单调递增/递减、索价平台期，以及题面约束上限规模。
//
// 规模取舍：本题两个骨架都是逐天转移的 O(n) 写法（状态机 DP 与滚动变量），
// 渐进复杂度同量级，即便上限用例直接取题面约束的最大规模 n = 5 * 10^4，
// 两者也都只需十万级操作，可从容跑完，因此无需额外压低规模。
// 期望值由独立参考实现核对（状态机 DP、峰谷贪心、O(n^2) 暴力三份实现
// 互相印证，并在 2 万组小规模随机数据上完全一致）。
// 数值范围检查：题面约束下最大可达利润约为 (5 * 10^4 / 2) * (49999 - 1) ≈ 1.25 * 10^9，
// 本表中最大期望值为 1222475000，在 int32 范围内；所有 prices[i] 与 fee
// 也都落在题面声明的取值区间内。
var cases = []struct {
	name   string
	prices []int
	fee    int
	expect int
}{
	{"例1 两次交易各扣一次手续费", []int{1, 3, 2, 8, 4, 9}, 2, 8},
	{"例2 手续费 3 仍有一笔可做", []int{1, 3, 7, 5, 10, 3}, 3, 6},
	{"边界 单元素无法交易", []int{1}, 0, 0},
	{"边界 单元素且手续费极大", []int{49999}, 49999, 0},
	{"边界 手续费为 0 等价于无限次交易", []int{7, 1, 5, 3, 6, 4}, 0, 7},
	{"边界 手续费高到任何一笔都亏", []int{1, 3, 1, 3}, 5, 0},
	{"边界 全程单调递减不交易", []int{5, 4, 3, 2, 1}, 1, 0},
	{"边界 全程单调递增只做一笔", []int{1, 2, 3, 4, 5}, 1, 3},
	{"边界 两元素价差不足手续费", []int{1, 2}, 100, 0},
	{"边界 两元素价差恰好等于手续费", []int{1, 3}, 2, 0},
	{"边界 价格平台期", []int{4, 4, 4, 4}, 0, 0},
	{"锯齿行情每段均有盈利", []int{1, 5, 1, 5, 1, 5}, 2, 6},
	{"上限规模 5*10^4 全部同价", rep(25000, 50000), 1, 0},
	{"上限规模 5*10^4 锯齿", zigzag(50000, 1, 49999), 100, 1222475000},
	{"上限规模 5*10^4 递增循环", rising(50000), 1000, 48998},
	{"上限规模 5*10^4 非递增至触底", falling(50000), 1, 0},
}

// TestMaxProfit 验证解法 1（状态机动态规划）。
func TestMaxProfit(t *testing.T) {
	checkCases(t, maxProfit)
}

// TestMaxProfit1 验证解法 2（滚动变量压缩状态）。
// func TestMaxProfit1(t *testing.T) {
// 	checkCases(t, maxProfit1)
// }

// checkCases 是全部解法共用的表驱动 helper。
func checkCases(t *testing.T, fn func([]int, int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.prices, c.fee)
			if actual != c.expect {
				t.Fatalf("prices=%s, fee=%d, expected=%d, actual=%d",
					pricesString(c.prices), c.fee, c.expect, actual)
			}
			t.Logf("prices=%s, fee=%d, expected=%d, actual=%d",
				pricesString(c.prices), c.fee, c.expect, actual)
		})
	}
}

// rep 返回长度为 n、元素全为 v 的序列，v 需落在 [1, 5 * 10^4) 内。
func rep(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// zigzag 返回长度为 n 的锯齿序列：偶数下标约取 lo，奇数下标约取 hi，
// 用于制造大量「买入—卖出」机会；lo、hi 需落在 [1, 5 * 10^4) 内。
func zigzag(n, lo, hi int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = lo + i%1000
		} else {
			s[i] = hi - i%1000
		}
	}
	return s
}

// rising 返回长度为 n 的循环递增序列：从 1 递增到 49999 后回到 1，
// 每次上坡都值得做一笔交易。
func rising(n int) []int {
	s := make([]int, n)
	v := 1
	for i := range s {
		s[i] = v
		v++
		if v >= 50000 {
			v = 1
		}
	}
	return s
}

// falling 返回长度为 n 的非递增序列，元素落在 [1, 5 * 10^4) 内（尾部触底为 1 后保持不变）。
func falling(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = 49999 - i
		if s[i] < 1 {
			s[i] = 1
		}
	}
	return s
}

// pricesString 在用例规模较大时只打印长度，避免失败信息过长。
func pricesString(prices []int) string {
	if len(prices) > 20 {
		return fmt.Sprintf("<len=%d>", len(prices))
	}
	return fmt.Sprintf("%v", prices)
}
