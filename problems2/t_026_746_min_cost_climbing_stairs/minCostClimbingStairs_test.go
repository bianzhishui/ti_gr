package t_026_746_min_cost_climbing_stairs

import "testing"

// cases 是全部解法共用的用例表，覆盖题面两个示例、两台阶/三台阶/全相等等边界，
// 以及约束上限规模（n = 1000，cost[i] = 999）。
//
// 规模取舍：本题两个骨架（滚动变量 DP、记忆化搜索）都只访问 n 个下标状态、
// 每个状态做常数次转移，渐进复杂度同为线性，因此上限规模用例直接取题面约束的
// 最大值 n = 1000，两个骨架都能在毫秒级内跑完，整个包远低于 1 秒目标。
//
// 期望值由独立参考实现核对：滚动变量 DP、记忆化搜索、朴素递归三份参考实现在
// 小规模随机数据（4000 组，n 取 2~10，cost[i] 取 0~999）上三方对齐，题面示例与
// 各用例答案再由参考实现逐一算出。
// 可能的最大答案为 500 * 999 = 499500（长度为 1000 时至少需要 500 次付费），
// 远小于 int 范围。
var cases = []struct {
	name   string
	cost   []int
	expect int
}{
	{"例1 三台阶从下标 1 起步", []int{10, 15, 20}, 15},
	{"例2 十台阶跳过所有 100", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	{"边界 两台阶均为 0", []int{0, 0}, 0},
	{"边界 两台阶取较小者", []int{1, 2}, 1},
	{"边界 两台阶前大后小", []int{5, 1}, 1},
	{"边界 两台阶都取到上限 999", []int{999, 999}, 999},
	{"三台阶递增", []int{1, 2, 3}, 2},
	{"三台阶递减", []int{3, 2, 1}, 2},
	{"三台阶中间极贵直接跨过", []int{1, 999, 1}, 2},
	{"六台阶费用全相等", []int{5, 5, 5, 5, 5, 5}, 15},
	{"前两级很贵后面便宜", []int{100, 100, 1, 1, 1, 1}, 102},
	{"固定 30 级序列", []int{3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2, 9, 16, 6, 13, 3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2}, 75},
	{"上限规模 n=1000 费用全为 999", rep(999, 1000), 499500},
	{"上限规模 n=1000 费用全为 0", make([]int, 1000), 0},
	{"上限规模 n=1000 费用全为 1", rep(1, 1000), 500},
	{"上限规模 n=1000 交替 999/0", alternating(999, 0, 1000), 0},
	{"上限规模 n=1000 交替 0/999", alternating(0, 999, 1000), 0},
}

// TestMinCostClimbingStairs 验证解法 1（动态规划 + 滚动变量）。
func TestMinCostClimbingStairs(t *testing.T) {
	checkCases(t, minCostClimbingStairs)
}

// TestMinCostClimbingStairs1 验证解法 2（记忆化搜索）。
func TestMinCostClimbingStairs1(t *testing.T) {
	checkCases(t, minCostClimbingStairs1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cost := append([]int(nil), c.cost...)
			actual := fn(cost)
			if actual != c.expect {
				t.Fatalf("cost=%v, expected=%v, actual=%v", c.cost, c.expect, actual)
			}
			t.Logf("cost=%v, expected=%v, actual=%v", c.cost, c.expect, actual)
		})
	}
}

// rep 返回 n 个相同元素组成的切片。
func rep(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// alternating 返回长度为 n、在 a 与 b 之间交替的切片。
func alternating(a, b, n int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = a
		} else {
			s[i] = b
		}
	}
	return s
}
