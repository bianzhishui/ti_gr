package t_009_115_distinct_subsequences

import (
	"strings"
	"testing"
)

// cases 覆盖题面全部示例、边界情况与上限规模，由全部解法共用。
// 期望值由独立参考实现算出
// （二维 DP 与滚动数组两份实现互相印证），且均不超过题面声明的 32 位有符号整数范围。
//
// 规模取舍说明：本题两个骨架的时间复杂度都是 O(|s|*|t|)，区别只在空间，
// 因此上限规模直接取题面约束的 |s|=|t|=1000（约 10^6 次状态转移），
// 两个骨架都能轻松跑完，整包远低于 1 秒。
// 另外，为保证答案不超出题面声明的 32 位范围，长 s 配短 t 的用例把 |t| 取成 3
// （C(1000,3)=166167000）；若取 |t|=4 则答案为 41417124750，已超出 32 位，故不采用。
var cases = []struct {
	name   string
	s      string
	t      string
	expect int
}{
	{"例1 rabbbit 中出现 rabbit", "rabbbit", "rabbit", 3},
	{"例2 babgbag 中出现 bag", "babgbag", "bag", 5},
	{"边界 单字符相同", "a", "a", 1},
	{"边界 单字符不同", "a", "b", 0},
	{"边界 t 比 s 长", "a", "aa", 0},
	{"边界 s 为空串", "", "a", 0},
	{"边界 t 为空串", "a", "", 1},
	{"边界 s 中无目标字符", "abc", "d", 0},
	{"边界 t 与 s 完全相同", "bag", "bag", 1},
	{"重复字符 三选二", "aaa", "aa", 3},
	{"重复字符 三选一", "aaa", "a", 3},
	{"交错重复字符", "babgbagg", "bag", 9},
	{"上限规模 1000x1000 全相同", strings.Repeat("a", 1000), strings.Repeat("a", 1000), 1},
	{"上限规模 1000x1000 字符不相交", strings.Repeat("a", 1000), strings.Repeat("b", 1000), 0},
	{"上限规模 1000x1000 交错串整体匹配", strings.Repeat("ab", 500), strings.Repeat("ab", 500), 1},
	{"上限规模 s 长 1000、t 为 3 个 a 的组合数", strings.Repeat("a", 1000), "aaa", 166167000},
}

// TestNumDistinct 验证解法 1（二维动态规划计数）。
func TestNumDistinct(t *testing.T) {
	checkCases(t, numDistinct)
}

// TestNumDistinct1 验证解法 2（滚动数组）。
func TestNumDistinct1(t *testing.T) {
	checkCases(t, numDistinct1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string, string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s, c.t)
			if actual != c.expect {
				t.Fatalf("s=%q, t=%q, expected=%v, actual=%v", c.s, c.t, c.expect, actual)
			}
			t.Logf("s=%q, t=%q, expected=%v, actual=%v", c.s, c.t, c.expect, actual)
		})
	}
}
