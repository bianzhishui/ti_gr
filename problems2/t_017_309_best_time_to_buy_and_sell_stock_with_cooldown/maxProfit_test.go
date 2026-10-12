package t_017_309_best_time_to_buy_and_sell_stock_with_cooldown

import (
	"fmt"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= prices.length <= 5000，两个骨架都是按天转移的 O(n) 状态机，
// 处于同一渐进量级，因此上限规模用例直接取到约束上限 5000，二者都能在 1 秒内从容跑完，
// 无需为其中某个骨架额外压低规模。
//
// 本文件所有 expected 均由独立参考实现（三状态数组 DP、滚动变量 DP、记忆化搜索三种写法
// 交叉验证，并另在随机小规模上与记忆化搜索逐一比对）核对得到，未凭记忆填写。
// 数值范围检查：prices[i] 均落在题面声明的 [0, 1000] 内，
// 最大答案出现在 5000 天 0/1000 交替的用例上（1250000），远未超出 int 范围。

// altSeq 返回长度为 n 的 0/1000 交替序列，覆盖价格在值域两端反复跳动的上限规模场景。
func altSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = 0
		} else {
			s[i] = 1000
		}
	}
	return s
}

// modSeq 返回长度为 n 的 i%1001 序列，覆盖缓慢抬高后归零的锯齿形态。
func modSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i % 1001
	}
	return s
}

// lcgSeq 用线性同余生成器返回长度为 n 的伪随机序列，元素落在 [0, 1000]，
// 用于在没有明显规律的价格序列上检验解法。
func lcgSeq(n int) []int {
	s := make([]int, n)
	x := int64(12345)
	for i := range s {
		x = (x*1103515245 + 12345) & 0x7fffffff
		s[i] = int(x % 1001)
	}
	return s
}

// cases 覆盖题面全部示例、边界情况与上限规模，由全部解法共用。
// 题面约束为 1 <= prices.length <= 5000，故不含空数组。
var cases = []struct {
	name     string
	prices   []int
	expected int
}{
	{"例1 买卖+冷冻", []int{1, 2, 3, 0, 2}, 3},
	{"例2 单元素", []int{1}, 0},
	{"单元素最小价格 0", []int{0}, 0},
	{"两元素递增", []int{1, 2}, 1},
	{"两元素递减", []int{2, 1}, 0},
	{"两元素相等", []int{1000, 1000}, 0},
	{"全部相等", []int{3, 3, 3, 3, 3}, 0},
	{"单调递增只做一笔最赚", []int{1, 2, 3, 4, 5}, 4},
	{"单调递减无交易可做", []int{5, 4, 3, 2, 1}, 0},
	{"值域两端先买后卖", []int{0, 1000}, 1000},
	{"值域两端先高后低", []int{1000, 0}, 0},
	{"冷冻期迫使放弃一次低买", []int{1, 5, 1, 5, 1, 5}, 8},
	{"多次低买高卖", []int{2, 1, 4, 5, 2, 9, 7}, 10},
	{"先跌后涨", []int{6, 1, 3, 2, 4, 7}, 6},
	{"含平台期", []int{2, 2, 5, 5, 1, 8}, 10},
	{"连涨时持有优于中途卖出", []int{1, 2, 4}, 3},
	{"约束上限 5000 天 0/1000 交替", altSeq(5000), 1250000},
	{"约束上限 5000 天 i%1001 锯齿", modSeq(5000), 4991},
	{"约束上限 5000 天伪随机序列", lcgSeq(5000), 657859},
}

// TestMaxProfit 验证解法 1（状态机动态规划）。
func TestMaxProfit(t *testing.T) {
	checkProfitCases(t, maxProfit)
}

// TestMaxProfit1 验证解法 2（滚动变量压缩状态）。
func TestMaxProfit1(t *testing.T) {
	checkProfitCases(t, maxProfit1)
}

func checkProfitCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.prices)
			if actual != c.expected {
				t.Fatalf("prices=%s, expected=%v, actual=%v", pricesString(c.prices), c.expected, actual)
			}
			t.Logf("prices=%s, expected=%v, actual=%v", pricesString(c.prices), c.expected, actual)
		})
	}
}

// pricesString 在用例规模较大时只打印长度，避免失败信息过长。
func pricesString(prices []int) string {
	if len(prices) > 20 {
		return fmt.Sprintf("<len=%d>", len(prices))
	}
	return fmt.Sprintf("%v", prices)
}
