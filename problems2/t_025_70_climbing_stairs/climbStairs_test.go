package t_025_70_climbing_stairs

import "testing"

// climbStairsCases 由全部解法共用，覆盖题面全部示例、n 的下边界与约束上限。
//
// 规模取舍：题面约束为 1 <= n <= 45，上限规模用例就取 n = 45（答案为 1836311903）。
// 递推骨架是 O(n)、矩阵快速幂骨架是 O(log n)，两者在 n = 45 这一规模下都是微秒级，
// 整包远在 1 秒预算之内，因此所有用例都直接取到约束上限，无需为某个骨架缩小规模。
//
// 期望值由独立参考实现核对（递推、2x2 矩阵快速幂、朴素递归三者交叉验证），
// 全部落在题面声明的 1 <= n <= 45 范围内，最大答值 1836311903 未超过 int32 上限 2147483647。
var climbStairsCases = []struct {
	name     string
	n        int
	expected int
}{
	{"例1 n=2", 2, 2},
	{"例2 n=3", 3, 3},
	{"边界 n=1", 1, 1},
	{"n=4", 4, 5},
	{"n=5", 5, 8},
	{"n=10", 10, 89},
	{"n=20", 20, 10946},
	{"约束上限 n=45", 45, 1836311903},
}

// TestClimbStairs 验证解法 1（递推 + 滚动变量）。
func TestClimbStairs(t *testing.T) {
	checkClimbStairsCases(t, climbStairs)
}

// TestClimbStairs1 验证解法 2（矩阵快速幂）。
func TestClimbStairs1(t *testing.T) {
	checkClimbStairsCases(t, climbStairs1)
}

// checkClimbStairsCases 用同一张用例表驱动单个解法。
func checkClimbStairsCases(t *testing.T, fn func(int) int) {
	t.Helper()
	for _, c := range climbStairsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.n)
			if actual != c.expected {
				t.Fatalf("n=%d, expected=%d, actual=%d", c.n, c.expected, actual)
			}
			t.Logf("n=%d, expected=%d, actual=%d", c.n, c.expected, actual)
		})
	}
}
