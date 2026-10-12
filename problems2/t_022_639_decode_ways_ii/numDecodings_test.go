package t_022_639_decode_ways_ii

import (
	"fmt"
	"strings"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= s.length <= 10^5，两个骨架分别是滚动数组线性动态规划与记忆化搜索，渐进复杂度都只与 n 相关，
// 不存在明显更差的骨架，因此上限规模用例直接取到约束上限 10^5 即可，两者都能在 1 秒内轻松跑完
// （记忆化搜索的递归深度虽为 10^5，Go 的栈可自动增长，实测仍在毫秒级）。
//
// cases 覆盖题面全部示例与边界情况，由全部解法共用。
// 期望值由独立参考实现核对：以滚动数组动态规划、记忆化搜索、完整 dp 数组三份实现交叉验证，
// 短串另用一份展开所有 '*' 取值的指数级回溯实现复核，结果完全一致。
// 数值范围检查通过：题面要求返回 10^9 + 7 的模，全部期望值均为 [0, 10^9+6] 区间内的非负整数。

// starBigCase 生成恰好 10^5 个 '*' 的上限规模输入：每一位都能独立成码或组合成码，答案在取模后仍非平凡。
func starBigCase() string {
	return strings.Repeat("*", 100000)
}

// mixBigCase 生成恰好 10^5 个字符的上限规模输入：'1*' 反复出现，覆盖数字与 '*' 的交替组合。
func mixBigCase() string {
	return strings.Repeat("1*", 50000)
}

var cases = []struct {
	name     string
	s        string
	expected int
}{
	{"例1 单个星号代表 1-9", "*", 9},
	{"例2 星号与前导 1 组合", "1*", 18},
	{"星号后接 7 只有 17 可配对", "*7", 10},
	{"例3 星号与前导 2 组合", "2*", 15},
	{"单元素普通数字", "1", 1},
	{"单元素 '0' 无法解码", "0", 0},
	{"两位数可作为整体", "12", 2},
	{"两个星号", "**", 96},
	{"前导 '0' 后接星号仍然无解", "0*", 0},
	{"星号后接 '0' 只能组合解码", "*0", 2},
	{"不含星号的经典解码参照", "226", 3},
	{"以 0 结尾但 0 无前导可配对", "100", 0},
	{"数字与星号交替出现", "1*1*1*", 6516},
	{"星号后接 27 超出字母映射", "*27", 11},
	{"多个零穿插的较长串", "1201234", 3},
	{"约束上限规模 10^5 全星号", starBigCase(), 81890275},
	{"约束上限规模 10^5 数字与星号交替", mixBigCase(), 860765661},
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
