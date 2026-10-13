package t_007_583_delete_operation_for_two_strings

import (
	"strings"
	"testing"
)

// cases 覆盖题面全部示例、边界情况与约束上限规模，由全部解法共用。
// 上限规模取题面声明的 500x500：两个骨架都是 O(m*n) 的二维 DP，约 25 万次转移，
// 整包可轻松跑完，无需为速度做额外取舍。
// 期望值由三份互相独立的参考实现
// （LCS 推导、直接二维 DP、0-1 BFS 最短路）一致算出；最大值 1000，远在 int 范围内。
var cases = []struct {
	name     string
	word1    string
	word2    string
	expected int
}{
	{"例1 sea 与 eat", "sea", "eat", 2},
	{"例2 leetcode 与 etco", "leetcode", "etco", 4},
	{"边界 单字符相同", "a", "a", 0},
	{"边界 单字符不同", "a", "b", 2},
	{"边界 两串完全相同", "abc", "abc", 0},
	{"边界 两串完全不同", "abc", "def", 6},
	{"边界 一个是另一个的子序列", "abc", "abcdefg", 4},
	{"边界 重复字符只保留较短公共部分", "aaa", "aa", 1},
	{"交错重复字符", "abcbdab", "bdcaba", 5},
	{"上限规模 500x500 全相同", strings.Repeat("a", 500), strings.Repeat("a", 500), 0},
	{"上限规模 500x500 完全不同", strings.Repeat("a", 500), strings.Repeat("b", 500), 1000},
	{"上限规模 500x500 周期串", strings.Repeat("ab", 250), strings.Repeat("ba", 250), 2},
	{"上限规模 500x500 错位周期串", strings.Repeat("abc", 166) + "ab", strings.Repeat("bca", 166) + "bc", 2},
}

// TestMinDistance 验证解法 1（先求 LCS）。
// func TestMinDistance(t *testing.T) {
// 	checkCases(t, minDistance)
// }

// TestMinDistance1 验证解法 2（直接二维 DP）。
func TestMinDistance1(t *testing.T) {
	checkCases(t, minDistance1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string, string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.word1, c.word2)
			if actual != c.expected {
				t.Fatalf("word1=%q, word2=%q, expected=%v, actual=%v", c.word1, c.word2, c.expected, actual)
			}
			t.Logf("word1=%q, word2=%q, expected=%v, actual=%v", c.word1, c.word2, c.expected, actual)
		})
	}
}
