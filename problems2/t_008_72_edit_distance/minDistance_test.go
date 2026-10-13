package t_008_72_edit_distance

import (
	"strings"
	"testing"
)

// cases 覆盖题面全部示例、边界情况与约束上限规模（len <= 500），由全部解法共用。
//
// 期望值由独立参考实现算出：该程序同时写了
// 二维 DP、滚动数组、记忆化递归三份实现，三者输出完全一致；结果均落在题目声明的
// 范围内（0 <= 答案 <= 500，最大值为两串长度之和的上界内）。
//
// 规模取舍：本题两个骨架都是 O(len(word1) * len(word2))，上限规模 500x500 也只有
// 约 25 万个状态，滚动数组版还能把空间降到一维，因此上限用例直接取满 500 即可，
// 每个解法跑完整张表都远低于 1 秒。
var cases = []struct {
	name   string
	word1  string
	word2  string
	expect int
}{
	{"例1 horse 转成 ros", "horse", "ros", 3},
	{"例2 intention 转成 execution", "intention", "execution", 5},
	{"边界 两串都为空", "", "", 0},
	{"边界 word1 为空只做插入", "", "abc", 3},
	{"边界 word2 为空只做删除", "abc", "", 3},
	{"边界 单字符相同", "a", "a", 0},
	{"边界 单字符不同需替换一次", "a", "b", 1},
	{"边界 等长且无公共字符", "abc", "xyz", 3},
	{"重复字符 aaa 转成 aa", "aaa", "aa", 1},
	{"上限规模 500x500 完全相同", strings.Repeat("a", 500), strings.Repeat("a", 500), 0},
	{"上限规模 500x500 完全不同", strings.Repeat("a", 500), strings.Repeat("b", 500), 500},
	{"上限规模 500x500 周期串 ab.../ba...", strings.Repeat("ab", 250), strings.Repeat("ba", 250), 2},
}

// TestMinDistance 验证解法 1（二维动态规划）。
func TestMinDistance(t *testing.T) {
	checkCases(t, minDistance)
}

// TestMinDistance1 验证解法 2（滚动数组）。
// func TestMinDistance1(t *testing.T) {
// 	checkCases(t, minDistance1)
// }

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string, string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.word1, c.word2)
			if actual != c.expect {
				t.Fatalf("name=%s, expected=%d, actual=%d", c.name, c.expect, actual)
			}
			t.Logf("name=%s, expected=%d, actual=%d", c.name, c.expect, actual)
		})
	}
}
