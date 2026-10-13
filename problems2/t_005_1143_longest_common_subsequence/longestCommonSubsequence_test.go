package t_005_1143_longest_common_subsequence

import (
	"strings"
	"testing"
)

// cases 覆盖题面全部示例与边界情况（含约束上限规模），由全部解法共用。
// 期望值由 /tmp 下独立参考程序（二维 DP 与滚动数组两份实现互相印证）算出。
var cases = []struct {
	name   string
	text1  string
	text2  string
	expect int
}{
	{"例1 abcde 与 ace", "abcde", "ace", 3},
	{"例2 完全相同", "abc", "abc", 3},
	{"例3 无公共字符", "abc", "def", 0},
	{"边界 单字符相同", "a", "a", 1},
	{"边界 单字符不同", "a", "b", 0},
	{"边界 无公共字符", "abc", "xyz", 0},
	{"边界 一个是另一个的子序列", "abc", "abcdefg", 3},
	{"边界 重复字符取较短公共部分", "aaa", "aa", 2},
	{"交错重复字符", "abcbdab", "bdcaba", 4},
	{"上限规模 1000x1000 全相同", strings.Repeat("a", 1000), strings.Repeat("a", 1000), 1000},
	{"上限规模 1000x1000 完全不同", strings.Repeat("a", 1000), strings.Repeat("b", 1000), 0},
	{"上限规模 1000x1000 周期串", strings.Repeat("ab", 500), strings.Repeat("ba", 500), 999},
}

// TestLongestCommonSubsequence 验证解法 1（二维动态规划）。
func TestLongestCommonSubsequence(t *testing.T) {
	checkCases(t, longestCommonSubsequence)
}

// TestLongestCommonSubsequence1 验证解法 2（滚动数组）。
// func TestLongestCommonSubsequence1(t *testing.T) {
// 	checkCases(t, longestCommonSubsequence1)
// }

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string, string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.text1, c.text2)
			if actual != c.expect {
				t.Fatalf("text1=%q, text2=%q, expected=%v, actual=%v", c.text1, c.text2, c.expect, actual)
			}
			t.Logf("text1=%q, text2=%q, expected=%v, actual=%v", c.text1, c.text2, c.expect, actual)
		})
	}
}
