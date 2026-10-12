package t_031_1039_minimum_score_triangulation_of_polygon

import "testing"

// cases 覆盖题面全部示例、边界情况与约束上限规模，由全部解法共用。
// 期望值由独立参考实现核对（自底向上区间 DP 与自顶向下记忆化搜索两份实现互相印证，
// 小规模用例另用暴力枚举再印证一次），不是凭记忆填写；
// 最大值 48000000 远小于 int 范围，也符合题面声明的数值范围。
//
// 规模取舍：两个骨架的时间复杂度都是 O(n^3)，n 取题面上限 50 时也只有约 1.25e5 次转移，
// 而且记忆化搜索的递归深度不超过 n，因此上限规模用例直接取满 n = 50，两个骨架都能从容跑完。
var cases = []struct {
	name   string
	values []int
	expect int
}{
	{"例1 三角形", []int{1, 2, 3}, 6},
	{"例2 四边形", []int{3, 7, 4, 5}, 144},
	{"例3 六边形", []int{1, 3, 1, 4, 1, 5}, 13},
	{"边界 最小三角形全1", []int{1, 1, 1}, 1},
	{"边界 最大三角形全100", []int{100, 100, 100}, 1000000},
	{"边界 正方形全1", []int{1, 1, 1, 1}, 2},
	{"边界 正方形交错取值", []int{1, 2, 1, 2}, 4},
	{"边界 五边形全1", []int{1, 1, 1, 1, 1}, 3},
	{"边界 五边形含大值", []int{1, 100, 1, 100, 1}, 201},
	{"六边形大小交替", []int{1, 100, 1, 100, 1, 100}, 301},
	{"十边形递增", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 328},
	{"上限规模 n=50 全1", filledValues(50, 1), 48},
	{"上限规模 n=50 全100", filledValues(50, 100), 48000000},
	{"上限规模 n=50 伪随机", []int{
		36, 87, 94, 99, 26, 94, 91, 43, 55, 96,
		36, 85, 52, 68, 35, 68, 6, 65, 89, 85,
		84, 46, 78, 98, 67, 19, 30, 94, 6, 69,
		76, 47, 79, 7, 16, 98, 20, 44, 32, 26,
		38, 69, 48, 100, 96, 45, 50, 90, 31, 46,
	}, 1020330},
}

// TestMinScoreTriangulation 验证解法 1（区间动态规划）。
func TestMinScoreTriangulation(t *testing.T) {
	checkCases(t, minScoreTriangulation)
}

// TestMinScoreTriangulation1 验证解法 2（记忆化搜索）。
func TestMinScoreTriangulation1(t *testing.T) {
	checkCases(t, minScoreTriangulation1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.values)
			if actual != c.expect {
				t.Fatalf("values=%v, expected=%v, actual=%v", c.values, c.expect, actual)
			}
			t.Logf("n=%d, expected=%v, actual=%v", len(c.values), c.expect, actual)
		})
	}
}

// filledValues 生成长度为 n、每个元素都等于 v 的切片。
func filledValues(n int, v int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = v
	}
	return values
}
