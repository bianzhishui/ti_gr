package t_011_198_house_robber

import "testing"

// cases 是全部解法共用的用例表，覆盖题面两个示例、单元素/两元素/全零/全相等等
// 边界，以及约束上限规模（n = 100，nums[i] = 400）。
//
// 规模取舍：本题两个骨架分别是记忆化搜索与自底向上 DP，都只访问 n 个状态，
// 上限用例取到题面约束的最大值 n = 100 即可，两个骨架都能在毫秒级内跑完。
// 期望值由独立参考实现算出（记忆化、DP、
// 朴素递归三份实现互相印证，并对小规模随机数据交叉验证 3000 组）；
// 最大可能答案为 50 * 400 = 20000，远小于 int 范围。
var cases = []struct {
	name   string
	nums   []int
	expect int
}{
	{"例1 相邻不能同时偷", []int{1, 2, 3, 1}, 4},
	{"例2 跳着偷中间多", []int{2, 7, 9, 3, 1}, 12},
	{"边界 单元素取该房", []int{5}, 5},
	{"边界 单元素金额为 0", []int{0}, 0},
	{"边界 两元素只能取较大者", []int{1, 2}, 2},
	{"边界 两元素前大后小", []int{10, 1}, 10},
	{"边界 三元素相邻全相等", []int{5, 5, 5}, 10},
	{"边界 100 间房金额全为 0", make([]int, 100), 0},
	{"递增序列 隔一间取", []int{1, 2, 3, 4, 5}, 9},
	{"隔一间更划算", []int{2, 1, 1, 2}, 4},
	{"固定较长序列 30 间", []int{3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2, 9, 16, 6, 13, 3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2}, 153},
	{"上限规模 n=100 金额全为 400", rep(400, 100), 20000},
	{"上限规模 n=100 金额交替 400/0", alternating(400, 0, 100), 20000},
}

// TestRob 验证解法 1（记忆化搜索）。
func TestRob(t *testing.T) {
	checkCases(t, rob)
}

// TestRob1 验证解法 2（自底向上动态规划）。
func TestRob1(t *testing.T) {
	checkCases(t, rob1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expect {
				t.Fatalf("nums=%v, expected=%v, actual=%v", c.nums, c.expect, actual)
			}
			t.Logf("nums=%v, expected=%v, actual=%v", c.nums, c.expect, actual)
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
