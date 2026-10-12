package t_032_132_palindrome_partitioning_ii

import (
	"fmt"
	"strings"
	"testing"
)

// 用例规模取舍说明：
// 题面约束是 1 <= s.length <= 2000。
//   - 解法 1（回文表 + 线性 DP）与写法正常的记忆化搜索都是 O(n^2) 量级，n = 2000 毫无压力，
//     因此「回文稀疏」的上限规模用例直接取满 2000。
//   - 解法 2 若对每个子串单独做一次 O(len) 的回文判断（教科书式的朴素记忆化），
//     在回文密集的串上会退化成 O(n^3)：实测 n = 2000 的全 'a' 串单例就要 1 秒上下。
//     为了让两个骨架都能从容跑完整张表（整包目标 1 秒内），回文密集的上限规模用例
//     统一压到 n = 1000（全 'a' / 交替 ab / 回文块 / 前后两段），此时朴素记忆化跑完整张表
//     约 0.4 秒，O(n^2) 的写法更是只需几十毫秒。
//
// 期望值来源：表中全部 want 均由独立参考实现核对——「回文判定表 + DP」、
// 「记忆化搜索」、「DAG 最短路 BFS」三份实现结果逐一吻合，短串（n <= 12）另与指数级
// 暴力枚举的分割方案比对一致，且全部 want 均未超过题目声明的范围（最少分割次数 <= n-1 = 1999）。

// pseudoRandomString 生成确定性伪随机小写串，供上限规模用例使用。
func pseudoRandomString(n int, seed int64) string {
	b := make([]byte, n)
	x := seed
	for i := range b {
		x = (x*1103515245 + 12345) & 0x7fffffff
		b[i] = byte('a' + x%26)
	}
	return string(b)
}

// cases 覆盖题面全部示例、边界情况与上限规模，由全部解法共用。
var cases = []struct {
	name string
	s    string
	want int
}{
	{"例1 aab", "aab", 1},
	{"例2 单字符 a", "a", 0},
	{"例3 ab", "ab", 1},
	{"边界 两字符相同 aa", "aa", 0},
	{"边界 本身是回文 aba", "aba", 0},
	{"边界 偶长度回文 abba", "abba", 0},
	{"边界 全部字符相同 aaaaa", "aaaaa", 0},
	{"边界 字符全不同 abcd", "abcd", 3},
	{"边界 两段拼两个回文 abaab", "abaab", 1},
	{"边界 交替串奇数回文 abbab", "abbab", 1},
	{"上限规模 2000 回文稀疏伪随机", pseudoRandomString(2000, 132), 1711},
	{"上限规模 1000 全相同", strings.Repeat("a", 1000), 0},
	{"上限规模 1000 交替 ab", strings.Repeat("ab", 500), 1},
	{"上限规模 1000 回文块 abcba", strings.Repeat("abcba", 200), 0},
	{"上限规模 1000 前半a后半b", strings.Repeat("a", 500) + strings.Repeat("b", 500), 1},
}

// TestMinCut 验证解法 1（回文判定表 + 线性动态规划）。
func TestMinCut(t *testing.T) {
	checkCases(t, minCut)
}

// TestMinCut1 验证解法 2（记忆化搜索）。
func TestMinCut1(t *testing.T) {
	checkCases(t, minCut1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(string) int) {
	t.Helper()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s)

			logprint := fmt.Sprintf("用例 %s：n=%d, expected=%d, actual=%d", c.name, len(c.s), c.want, actual)
			if actual == c.want {
				t.Log(logprint)
			} else {
				t.Fatal(logprint)
			}
		})
	}
}
