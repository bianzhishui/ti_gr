package t_020_188_best_time_to_buy_and_sell_stock_iv

import (
	"fmt"
	"testing"
)

// 复杂度取舍说明：本题两个骨架（三维 DP、滚动数组压缩天数）的渐进复杂度都是
// O(n × k)（n = len(prices)），都能从容处理题目约束的上限，因此上限规模用例
// 直接取满 n = 1000、k = 100，整张表跑完远低于 1 秒，无需为某个骨架裁减规模。
//
// 期望值来源：表中所有 want 均由独立参考实现核对（三维 DP / 记忆化搜索 /
// k 饱和时的无限次交易贪心三种写法互相验证一致），并确认不超过题目声明的
// 数值范围（k <= 100、prices[i] <= 1000，故答案上界为 100 × 1000 = 100000）。

// largePrices 生成固定的伪随机价格序列，n = 1000，取值 0..1000。
func largePrices() []int {
	prices := make([]int, 1000)
	x := 12345
	for i := range prices {
		x = (x*1103515245 + 12345) & 0x7fffffff
		prices[i] = x % 1001
	}
	return prices
}

// sawtooth 生成 0,1000,0,1000,... 的锯齿序列，n = 1000。
func sawtooth() []int {
	prices := make([]int, 1000)
	for i := range prices {
		if i%2 == 0 {
			prices[i] = 0
		} else {
			prices[i] = 1000
		}
	}
	return prices
}

// cases 由全部解法共用。
var cases = []struct {
	name     string
	k        int
	prices   []int
	expected int
}{
	{"例1 两笔交易只做其中一笔", 2, []int{2, 4, 1}, 2},
	{"例2 两段行情各赚一笔", 2, []int{3, 2, 6, 5, 0, 3}, 7},
	{"边界 单元素无法交易", 1, []int{5}, 0},
	{"边界 两元素上涨", 1, []int{1, 9}, 8},
	{"边界 两元素下跌", 1, []int{9, 1}, 0},
	{"边界 单调不降只赚一段", 2, []int{1, 2, 3, 4, 5}, 4},
	{"边界 单调不增无利可图", 2, []int{7, 6, 4, 3, 1}, 0},
	{"边界 价格全部相等", 3, []int{5, 5, 5, 5}, 0},
	{"边界 k=1 退化为只做一笔", 1, []int{7, 1, 5, 3, 6, 4}, 5},
	{"边界 k 大于天数退化为无限次交易", 100, []int{3, 3, 5, 0, 0, 3, 1, 4}, 8},
	{"上限 n=1000 价格全为零", 100, make([]int, 1000), 0},
	{"上限 n=1000 k=1 锯齿", 1, sawtooth(), 1000},
	{"上限 n=1000 k=100 锯齿", 100, sawtooth(), 100000},
	{"上限 n=1000 k=100 伪随机", 100, largePrices(), 85285},
}

// checkCases 让各解法跑同一张用例表。
func checkCases(t *testing.T, fn func(int, []int) int) {
	t.Helper()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.k, c.prices)

			logprint := fmt.Sprintf("k=%d, n=%d, expected=%d, actual=%d", c.k, len(c.prices), c.expected, actual)
			if actual == c.expected {
				t.Log(logprint)
			} else {
				t.Fatal(logprint)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) { checkCases(t, maxProfit) }

func TestMaxProfit1(t *testing.T) { checkCases(t, maxProfit1) }
