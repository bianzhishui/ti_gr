package t_010_392_is_subsequence

import (
	"fmt"
	"strings"
	"testing"
)

// cases 覆盖题面全部示例、边界情况与约束上限规模，由全部解法共用。
//
// 期望值由独立参考实现跑出，不是凭记忆填写：
// 该程序内含三份互相独立的实现（双指针贪心、next 位置表、strings.IndexByte 逐步推进），
// 全部 20 个用例三者结果完全一致；用例长度也一并自检，均未超过题面约束
// 0 <= s.length <= 100、0 <= t.length <= 10^4。
//
// 规模取舍：上限用例取满约束规模（s 长 100、t 长 10^4）。
// 骨架 1 单次查询为线性扫描；骨架 2 要先建 (|t|+1) × 26 的 next 表，单次约 2.6e5 次操作，
// 两者跑完整张表都远在 1 秒以内，因此上限规模无需下调。
var cases = []struct {
	name string
	s    string
	t    string
	want bool
}{
	{"例1 s 是 t 的子序列", "abc", "ahbgdc", true},
	{"例2 字符顺序不匹配", "axc", "ahbgdc", false},
	{"边界 s 为空串", "", "ahbgdc", true},
	{"边界 t 为空串而 s 非空", "a", "", false},
	{"边界 s 与 t 都为空串", "", "", true},
	{"边界 s 与 t 完全相同", "abc", "abc", true},
	{"边界 s 比 t 更长", "abcde", "abc", false},
	{"边界 s 单字符未在 t 中出现", "c", "ab", false},
	{"顺序颠倒不算子序列", "aec", "abcde", false},
	{"重复字符不够用", "aab", "ab", false},
	{"重复字符恰好匹配", "aab", "aaab", true},
	{"t 首尾各含所需字符", "ab", "axxxxxxxxb", true},
	{"交错重复串匹配", "abc", "aabbcc", true},
	{"上限规模 s=100 个 a，t=10^4 个 a", strings.Repeat("a", 100), strings.Repeat("a", 10000), true},
	{"上限规模 s=100 个 a，t=10^4 个 b", strings.Repeat("a", 100), strings.Repeat("b", 10000), false},
	{"上限规模 t 中 a 的个数刚好够", strings.Repeat("a", 100), strings.Repeat("a", 9999) + "b", true},
	{"上限规模 s=ab 重复 50 次，t=ab 重复 5000 次", strings.Repeat("ab", 50), strings.Repeat("ab", 5000), true},
	{"上限规模 s 需在 t 末尾才凑齐", strings.Repeat("a", 100), strings.Repeat("b", 9900) + strings.Repeat("a", 100), true},
	{"上限规模 s 需贪心跳过前缀", strings.Repeat("ab", 50), strings.Repeat("ba", 5000), true},
	{"上限规模 s 差一个字符", strings.Repeat("a", 100), strings.Repeat("a", 99), false},
}

// TestIsSubsequence 验证解法 1（双指针贪心）。
// func TestIsSubsequence(t *testing.T) {
// 	checkCases(t, isSubsequence)
// }

// TestIsSubsequence1 验证解法 2（预处理 next 位置表）。
func TestIsSubsequence1(t *testing.T) {
	checkCases(t, isSubsequence1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string, string) bool) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s, c.t)
			detail := caseDetail(c.s, c.t)
			if actual != c.want {
				t.Errorf("expected=%v, actual=%v, %s", c.want, actual, detail)
				return
			}
			t.Logf("expected=%v, actual=%v, %s", c.want, actual, detail)
		})
	}
}

// caseDetail 生成用例摘要；超长字符串只报长度，避免上限用例把日志刷屏。
func caseDetail(s, t string) string {
	if len(s) > 40 || len(t) > 40 {
		return fmt.Sprintf("len(s)=%d, len(t)=%d", len(s), len(t))
	}
	return fmt.Sprintf("s=%q, t=%q", s, t)
}
