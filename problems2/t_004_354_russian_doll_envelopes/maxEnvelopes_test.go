package t_004_354_russian_doll_envelopes

import "testing"

// cases 覆盖题面全部示例、边界情况（空输入、单信封、同宽 / 同高无法嵌套）与约束上限规模，
// 由全部解法共用。
// 注意：「规模上限 100000 严格递增链」含 100000 个信封，O(n^2) 的动态规划解法在该用例上约需 6 秒。
var cases = []struct {
	name      string
	envelopes [][]int
	expect    int
}{
	{"例1 标准套娃", [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
	{"例2 全部相同", [][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	{"空输入", [][]int{}, 0},
	{"单信封", [][]int{{1, 1}}, 1},
	{"同宽不同高不能嵌套", [][]int{{4, 5}, {4, 6}, {4, 7}}, 1},
	{"同高不同宽不能嵌套", [][]int{{2, 5}, {3, 5}, {4, 5}}, 1},
	{"乱序链需正确排序", [][]int{{6, 7}, {2, 3}, {5, 4}, {4, 5}, {1, 1}}, 4},
	{"宽高同步递增", [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}, 4},
	{"同宽高度降序消除伪递增", [][]int{{1, 1}, {2, 2}, {2, 3}, {3, 3}}, 3},
	{"宽高取值上限", [][]int{{100000, 100000}, {1, 1}}, 2},
	{"规模上限 100000 严格递增链", envelopeChain(100000), 100000},
	{"大规模 20000 全相同", envelopeRepeat(20000), 1},
}

// TestMaxEnvelopes 验证解法 1（排序 + 贪心 + 二分求最长严格递增子序列）。
func TestMaxEnvelopes(t *testing.T) {
	checkCases(t, maxEnvelopes)
}

// TestMaxEnvelopes1 验证解法 2（排序 + O(n^2) 动态规划）。
func TestMaxEnvelopes1(t *testing.T) {
	checkCases(t, maxEnvelopes1)
}

// checkCases 用共享的 cases 表驱动单个解法。
func checkCases(t *testing.T, fn func([][]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.envelopes)
			if actual != c.expect {
				t.Fatalf("case %s: expected=%v, actual=%v", c.name, c.expect, actual)
			}
			t.Logf("case %s: expected=%v, actual=%v", c.name, c.expect, actual)
		})
	}
}

// envelopeChain 生成 size 个宽高同步严格递增的信封：[1,1], [2,2], ..., [size,size]。
// 用于构造约束上限规模的用例（size = 100000，宽高均不超过 100000）。
func envelopeChain(size int) [][]int {
	envelopes := make([][]int, size)
	for i := range envelopes {
		envelopes[i] = []int{i + 1, i + 1}
	}
	return envelopes
}

// envelopeRepeat 生成 size 个完全相同的信封（宽高均为 1），用于验证同宽同高只能取一个。
func envelopeRepeat(size int) [][]int {
	envelopes := make([][]int, size)
	for i := range envelopes {
		envelopes[i] = []int{1, 1}
	}
	return envelopes
}
