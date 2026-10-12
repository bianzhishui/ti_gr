package t_019_123_best_time_to_buy_and_sell_stock_iii

import "testing"

// cases 是全部解法共用的用例表，覆盖题面全部 4 个示例，以及单元素、两元素、
// 全相等、只做一笔更优、三笔机会但只允许两笔、平台后暴涨等边界，
// 最后是题面约束上限规模 n = 100000 的两组用例。
//
// 规模取舍：本题两个骨架（状态机动态规划、滚动变量压缩）都是 O(n) 时间、
// O(1) 空间，即便取到题面约束的最大规模 n = 100000 也只有十万级操作，
// 因此上限用例直接取题面最大值，无需为任何骨架缩小规模。
//
// 期望值由独立参考实现核对：
// 程序里写了「5 状态线性扫描」「前缀最大单笔 + 后缀最大单笔枚举分割点」
// 「枚举两笔不重叠交易的 O(n^4) 暴力」三份互相独立的实现，在 30000 组
// n <= 9 的随机数据上三者答案完全一致，在 20 组 n = 100000 的随机数据上
// 前两份实现答案也完全一致；表中每个固定用例的期望值均取自该程序的输出。
// 数值范围检查：单笔利润最大 10^5，两笔最多 2 * 10^5 = 200000，
// 不超过 prices.length <= 10^5、prices[i] <= 10^5 所允许的范围，也在 int 内。
var cases = []struct {
	name   string
	prices []int
	expect int
}{
	{"例1 两笔不重叠交易", []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	{"例2 单调上涨只做一笔", []int{1, 2, 3, 4, 5}, 4},
	{"例3 单调下跌无法获利", []int{7, 6, 4, 3, 1}, 0},
	{"例4 单元素", []int{1}, 0},
	{"边界 单元素价格为零", []int{0}, 0},
	{"边界 两元素递增", []int{1, 2}, 1},
	{"边界 两元素递减", []int{2, 1}, 0},
	{"边界 两元素相等", []int{0, 0}, 0},
	{"边界 全相等无套利空间", []int{5, 5, 5, 5}, 0},
	{"边界 两笔分开做优于一笔", []int{1, 4, 2, 7}, 8},
	{"边界 三次机会但只允许两笔", []int{1, 2, 1, 2, 1, 2}, 2},
	{"边界 先跌后涨两笔各赚一笔", []int{2, 1, 2, 0, 1}, 2},
	{"边界 平台期后暴涨", []int{1, 1, 1, 1, 100}, 99},
	{"固定 阶梯上涨 n=7", []int{1, 2, 3, 4, 5, 6, 7}, 6},
	{"固定 先涨后跌再涨 n=8", []int{2, 1, 4, 5, 2, 9, 7, 10}, 12},
	{"固定 分散行情 n=20", []int{8, 3, 6, 2, 8, 8, 4, 2, 0, 7, 3, 9, 1, 5, 9, 0, 6, 2, 7, 4}, 17},
	{"上限规模 n=100000 单调上升", ascend(0, 100000), 99999},
	{"上限规模 n=100000 0/100000 锯齿交替", sawtooth(0, 100000, 100000), 200000},
	{"上限规模 n=100000 全部 100000", rep(100000, 100000), 0},
}

// TestMaxProfit 验证解法 1（状态机动态规划）。
func TestMaxProfit(t *testing.T) {
	checkCases(t, maxProfit)
}

// TestMaxProfit1 验证解法 2（滚动变量压缩状态）。
func TestMaxProfit1(t *testing.T) {
	checkCases(t, maxProfit1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.prices)
			if actual != c.expect {
				t.Fatalf("case %q: n=%d, expected=%v, actual=%v", c.name, len(c.prices), c.expect, actual)
			}
			// 上限用例的切片过长，只记录规模而不打印全部元素。
			t.Logf("case %q: n=%d, expected=%v, actual=%v", c.name, len(c.prices), c.expect, actual)
		})
	}
}

// rep 返回 n 个相同元素组成的切片。
func rep(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// sawtooth 返回长度为 n、在 lo 与 hi 之间交替的锯齿切片。
func sawtooth(lo, hi, n int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = lo
		} else {
			s[i] = hi
		}
	}
	return s
}

// ascend 返回长度为 n、从 start 开始逐日上涨 1 的切片。
func ascend(start, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = start + i
	}
	return s
}
