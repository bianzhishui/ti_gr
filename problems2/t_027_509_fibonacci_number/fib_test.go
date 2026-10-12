package t_027_509_fibonacci_number

import "testing"

// fibCases 由全部解法共用，覆盖题面全部示例、n 的下边界与约束上限。
//
// 规模取舍：题面约束为 0 <= n <= 30，上限规模用例就取 n = 30（F(30) = 832040）。
// 递推骨架与矩阵快速幂骨架在该规模下都是微秒级，整包远在 1 秒预算之内，无需额外缩小规模。
//
// 期望值由独立参考实现（递推、2x2 矩阵快速幂、朴素递归三者交叉验证）核对后填入，
// 全部落在题面声明的 0 <= n <= 30 范围内，答值 832040 远小于 int 上限。
var fibCases = []struct {
	name     string
	n        int
	expected int
}{
	{"例1 n=2", 2, 1},
	{"例2 n=3", 3, 2},
	{"例3 n=4", 4, 3},
	{"边界 n=0", 0, 0},
	{"边界 n=1", 1, 1},
	{"n=5", 5, 5},
	{"n=10", 10, 55},
	{"n=20", 20, 6765},
	{"约束上限 n=30", 30, 832040},
}

// TestFib 验证解法 1（递推 + 滚动变量）。
func TestFib(t *testing.T) {
	checkFibCases(t, fib)
}

// TestFib1 验证解法 2（矩阵快速幂）。
func TestFib1(t *testing.T) {
	checkFibCases(t, fib1)
}

// checkFibCases 用同一张用例表驱动单个解法。
func checkFibCases(t *testing.T, fn func(int) int) {
	t.Helper()
	for _, c := range fibCases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.n)
			if actual != c.expected {
				t.Fatalf("n=%d, expected=%d, actual=%d", c.n, c.expected, actual)
			}
			t.Logf("n=%d, expected=%d, actual=%d", c.n, c.expected, actual)
		})
	}
}
