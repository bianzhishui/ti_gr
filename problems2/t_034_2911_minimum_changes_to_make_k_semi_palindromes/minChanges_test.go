package t_034_2911_minimum_changes_to_make_k_semi_palindromes

import (
	"fmt"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 2 <= s.length <= 200、1 <= k <= s.length / 2。两个骨架（预处理子串代价 + 划分型
// 动态规划、预处理子串代价 + 记忆化搜索）的状态规模都由 n 与 k 决定，把规模取到约束上限
// n = 200（含 k 取上限 n/2）时，二者都只需几十毫秒，远低于 1 秒；记忆化搜索若多一层以
// (起点, 剩余段数) 为键的查表开销，同样不会超过时间上限。因此上限规模用例直接取
// n = 200 / 199，不再额外压低规模。
//
// cases 覆盖题面全部示例与边界情况（最短串、单段、k 取到 n/2 上限、全同字符、周期串、
// 奇数长度等），由全部解法共用。expected 均由独立参考实现核对：两套「子串改造成半回文串的
// 最少修改次数」实现（组内双指针比对 / 先拼出每组再判回文）、自底向上的划分型 DP、自顶向下的
// 记忆化搜索，以及穷举全部切分位置的暴力实现；四者在随机小规模串上全量比对一致，大规模用例上
// DP 与记忆化搜索亦逐例一致。数值范围检查通过：每段长度至少 2、每段最多改 floor(len/2) 次，
// 答案不会超过 n/2 = 100，远未超出 int 的表示范围。

var cases = []struct {
	name     string
	s        string
	k        int
	expected int
}{
	{"例1 s=abcac k=2", "abcac", 2, 1},
	{"例2 s=abcdef k=2", "abcdef", 2, 2},
	{"例3 s=aabbaa k=3", "aabbaa", 3, 0},
	{"最短串 aa k=1 已是半回文串", "aa", 1, 0},
	{"最短串 ab k=1 需改 1 次", "ab", 1, 1},
	{"abab k=1 取 d=2 已半回文", "abab", 1, 0},
	{"abab k=2 只能切成长度 2 的两段", "abab", 2, 2},
	{"aabb k=2 切成 aa 与 bb", "aabb", 2, 0},
	{"abcabc k=1 题面示例串", "abcabc", 1, 0},
	{"abcabc k=3 每段改 1 次", "abcabc", 3, 3},
	{"aaaaaa k=3 全同字符", "aaaaaa", 3, 0},
	{"上限规模 n=200 k=1 伪随机", randStr(200, 2911), 1, 72},
	{"上限规模 n=200 k=2 伪随机", randStr(200, 2911), 2, 58},
	{"上限规模 n=200 k=50 伪随机", randStr(200, 2911), 50, 45},
	{"上限规模 n=200 k=99 伪随机", randStr(200, 2911), 99, 92},
	{"上限规模 n=200 k=100 伪随机", randStr(200, 2911), 100, 95},
	{"上限规模 n=199 k=99 伪随机", randStr(199, 7), 99, 93},
	{"上限规模 n=200 k=1 全 a", repeatChar(200, 'a'), 1, 0},
	{"上限规模 n=200 k=100 全 a", repeatChar(200, 'a'), 100, 0},
	{"上限规模 n=200 k=1 周期 ab", cyclePattern("ab", 200), 1, 0},
	{"上限规模 n=200 k=100 周期 ab", cyclePattern("ab", 200), 100, 100},
}

// TestMinChanges 验证解法 1（预处理子串代价 + 划分型动态规划）。
func TestMinChanges(t *testing.T) {
	checkMinChangesCases(t, minChanges)
}

// TestMinChanges1 验证解法 2（预处理子串代价 + 记忆化搜索）。
func TestMinChanges1(t *testing.T) {
	checkMinChangesCases(t, minChanges1)
}

func checkMinChangesCases(t *testing.T, fn func(string, int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s, c.k)
			if actual != c.expected {
				t.Fatalf("s=%s, k=%d, expected=%d, actual=%d", sString(c.s), c.k, c.expected, actual)
			}
			t.Logf("s=%s, k=%d, expected=%d, actual=%d", sString(c.s), c.k, c.expected, actual)
		})
	}
}

// sString 在用例规模较大时只打印长度，避免日志与失败信息过长。
func sString(s string) string {
	if len(s) > 20 {
		return fmt.Sprintf("<len=%d>", len(s))
	}
	return fmt.Sprintf("%q", s)
}

// repeatChar 返回由 n 个字符 c 组成的字符串。
func repeatChar(n int, c byte) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = c
	}
	return string(b)
}

// cyclePattern 返回把 pattern 循环填充到长度 n 的字符串。
func cyclePattern(pattern string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = pattern[i%len(pattern)]
	}
	return string(b)
}

// randStr 用带固定种子的 xorshift32 生成伪随机小写字符串，保证用例可复现且贴近一般情况。
func randStr(n int, seed uint32) string {
	x := seed
	b := make([]byte, n)
	for i := range b {
		x ^= x << 13
		x ^= x >> 17
		x ^= x << 5
		b[i] = byte('a' + x%26)
	}
	return string(b)
}
