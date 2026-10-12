package t_021_91_decode_ways

import (
	"fmt"
	"strings"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= s.length <= 100，两个骨架（线性动态规划、记忆化搜索）的渐进复杂度都只与 n 相关，
// 上限规模用例取到约束上限 100 时两者都能在 1 秒内轻松跑完，因此不需要压低规模。
// 但题面声明「答案肯定是一个 32 位整数」，而长度为 100 的全 '1' 串其解数远超 32 位，
// 所以上限规模用例特意构造成答案为 2178309（远小于 2^31-1），避免出现超出题目答案范围的输入。

// bigCase 返回长度恰为 100 的上限规模用例：前 30 个 '1' 提供指数级增长，
// 后 35 组 "27" 因两数 27 超出 [10, 26] 而只能逐位单独解码，从而把答案压制在 32 位整数范围内。
func bigCase() string {
	return strings.Repeat("1", 30) + strings.Repeat("27", 35)
}

// cases 覆盖题面全部示例与边界情况，由全部解法共用。
// 期望值由独立参考实现核对：以线性动态规划、记忆化搜索、短串指数回溯三份实现交叉验证，
// 另用一套独立编写的迭代实现复核，结果完全一致；
// 数值范围检查通过：全部答案均为非负，最大值为上限规模用例的 2178309，未超出题面声明的 32 位整数范围。
var cases = []struct {
	name     string
	s        string
	expected int
}{
	{"例1 可作为整体或拆成两位", "12", 2},
	{"例2 两种分组方式叠加", "226", 3},
	{"例3 前导零导致无法解码", "06", 0},
	{"题面说明中的 11106", "11106", 2},
	{"单元素可解码", "1", 1},
	{"单元素 '0' 无法单独解码", "0", 0},
	{"两位数恰为 10", "10", 1},
	{"两位数 27 超出字母映射范围", "27", 1},
	{"以 0 结尾但 0 无前导可配对", "100", 0},
	{"中间夹 0 仍可整体解码", "101", 1},
	{"0 前面是 3，30 不是合法编码", "301", 0},
	{"含 0 的两位数与单数字交替", "2101", 1},
	{"多个零穿插的较长串", "1201234", 3},
	{"连续 9 无法组合，只有一种拆法", "9999999999", 1},
	{"约束上限规模 100，答案仍落在 32 位内", bigCase(), 2178309},
}

// TestNumDecodings 验证解法 1（线性动态规划）。
func TestNumDecodings(t *testing.T) {
	checkNumDecodingsCases(t, numDecodings)
}

// TestNumDecodings1 验证解法 2（记忆化搜索）。
func TestNumDecodings1(t *testing.T) {
	checkNumDecodingsCases(t, numDecodings1)
}

// checkNumDecodingsCases 用同一张 cases 表驱动一个解法骨架。
func checkNumDecodingsCases(t *testing.T, fn func(string) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.s)
			if actual != c.expected {
				t.Fatalf("s=%s, expected=%v, actual=%v", sString(c.s), c.expected, actual)
			}
			t.Logf("s=%s, expected=%v, actual=%v", sString(c.s), c.expected, actual)
		})
	}
}

// sString 在用例较长时只打印长度与前缀，避免失败信息过长。
func sString(s string) string {
	if len(s) > 20 {
		return fmt.Sprintf("<len=%d, prefix=%q>", len(s), s[:10])
	}
	return fmt.Sprintf("%q", s)
}
