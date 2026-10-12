package t_030_516_longest_palindromic_subsequence

import (
	"strings"
	"testing"
)

// cases 覆盖题面全部示例、边界情况与约束上限规模，由全部解法共用。
//
// 期望值由独立参考实现核对：区间 DP、反转串 LCS、记忆化递归三份实现在短串上
// 与暴力枚举全部子序列的结果逐一吻合，并在 1000 组随机串上互相印证。
//
// 两个骨架的渐进复杂度同为 O(n^2)，题面约束上限 s.length <= 1000 即为最坏规模
// （约 10^6 次状态转移），因此上限规模用例直接取到 1000，两个骨架都能从容跑完。
var cases = []struct {
	name   string
	s      string
	expect int
}{
	{"例1 bbbab", "bbbab", 4},
	{"例2 cbbd", "cbbd", 2},
	{"边界 单字符", "a", 1},
	{"边界 单字符最大字母", "z", 1},
	{"边界 两字符相同", "aa", 2},
	{"边界 两字符不同", "ab", 1},
	{"边界 全部字符相同", "aaaa", 4},
	{"边界 全部字符互不相同", "abcd", 1},
	{"边界 本身即回文", "abcba", 5},
	{"边界 偶长度回文", "abccba", 6},
	{"部分重复 前半a后半b", "aaaaabbbbb", 5},
	{"交错重复 abcabcabc", "abcabcabc", 5},
	{"嵌套重复 aaabaa", "aaabaa", 5},
	{"整体回文且含重复 aabbaa", "aabbaa", 6},
	{"周期串 abcdefabcdef", "abcdefabcdef", 3},
	{"上限规模 1000 全相同", strings.Repeat("a", 1000), 1000},
	{"上限规模 1000 交替 ab", strings.Repeat("ab", 500), 999},
	{"上限规模 1000 前半a后半b", strings.Repeat("a", 500) + strings.Repeat("b", 500), 500},
	{"上限规模 1000 周期 abcdefghij", strings.Repeat("abcdefghij", 100), 199},
}

// TestLongestPalindromeSubseq 验证解法 1（区间动态规划）。
func TestLongestPalindromeSubseq(t *testing.T) {
	checkCases(t, longestPalindromeSubseq)
}

// TestLongestPalindromeSubseq1 验证解法 2（原串与反转串的最长公共子序列）。
func TestLongestPalindromeSubseq1(t *testing.T) {
	checkCases(t, longestPalindromeSubseq1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s)
			if actual != c.expect {
				t.Fatalf("s=%q, expected=%v, actual=%v", c.s, c.expect, actual)
			}
			t.Logf("s=%q, expected=%v, actual=%v", c.s, c.expect, actual)
		})
	}
}
