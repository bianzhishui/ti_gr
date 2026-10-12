package t_024_139_word_break

import (
	"fmt"
	"strings"
	"testing"
)

// wordBreakCases 覆盖题面全部示例、边界情况与约束上限，由全部解法共用。
// 期望值由独立参考实现核对（自底向上 DP 与记忆化搜索两份实现结果一致），非凭记忆填写。
//
// 规模取舍：两个骨架都是多项式级别（约 O(n^2) 次子串查找），没有渐进复杂度明显更差的思路，
// 故上限规模用例直接取满题面约束 s.length = 300、wordDict.length = 1000，
// 两个骨架都能从容跑完，整个包的测试耗时远低于 1 秒。
var wordBreakCases = []struct {
	name     string
	s        string
	wordDict []string
	expected bool
}{
	{"例1 拆成两个单词", "leetcode", []string{"leet", "code"}, true},
	{"例2 单词可以重复使用", "applepenapple", []string{"apple", "pen"}, true},
	{"例3 字典无法覆盖全部字符", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	{"单字符命中字典", "a", []string{"a"}, true},
	{"单字符不在字典中", "a", []string{"b"}, false},
	{"字典单词长于目标串", "a", []string{"ab"}, false},
	{"整串即字典中的一个单词", "abcdefghij", []string{"abcdefghij"}, true},
	{"同一单词重复拼接", "aaaaaa", []string{"aa"}, true},
	{"需回溯选择分割点", "goalspecial", []string{"go", "goal", "goals", "special"}, true},
	{"结尾字符无法覆盖", "aaaaab", []string{"a", "aa", "aaa", "aaaa"}, false},
	{"约束上限 300 字符且可拆", strings.Repeat("ab", 150), buildDict("a", "b", "ab"), true},
	{"约束上限 300 字符且不可拆", strings.Repeat("ab", 149) + "az", buildDict("a", "b", "ab"), false},
}

// buildDict 造一个恰好 1000 个互不相同单词的字典：前几个用 base，
// 其余填充为形如 "q????" 的 5 字母单词，仅用于把 wordDict.length 撑到约束上限。
// 填充单词都以 'q' 开头，与用例中的 s 不可能匹配，不影响拆分结果。
func buildDict(base ...string) []string {
	dict := make([]string, 0, 1000)
	dict = append(dict, base...)
	for i := 0; len(dict) < 1000; i++ {
		w := []byte{'q', 'a', 'a', 'a', 'a'}
		n := i
		for j := 4; j >= 1; j-- {
			w[j] = byte('a' + n%26)
			n /= 26
		}
		dict = append(dict, string(w))
	}
	return dict
}

// dictDesc 让失败信息保持简短：字典很长时只展示前几个单词。
func dictDesc(wordDict []string) string {
	const maxShow = 6
	if len(wordDict) <= maxShow {
		return fmt.Sprintf("%v", wordDict)
	}
	return fmt.Sprintf("%v...（共 %d 个单词）", wordDict[:maxShow], len(wordDict))
}

// TestWordBreak 验证解法 1（动态规划）。
func TestWordBreak(t *testing.T) {
	checkWordBreakCases(t, wordBreak)
}

// TestWordBreak1 验证解法 2（记忆化搜索）。
func TestWordBreak1(t *testing.T) {
	checkWordBreakCases(t, wordBreak1)
}

func checkWordBreakCases(t *testing.T, fn func(string, []string) bool) {
	t.Helper()
	for _, c := range wordBreakCases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s, c.wordDict)
			if actual != c.expected {
				t.Fatalf("s=%q, wordDict=%s, expected=%v, actual=%v", c.s, dictDesc(c.wordDict), c.expected, actual)
			}
			t.Logf("s=%q, wordDict=%s, expected=%v, actual=%v", c.s, dictDesc(c.wordDict), c.expected, actual)
		})
	}
}
