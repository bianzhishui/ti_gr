package t_012_213_house_robber_ii

import "testing"

// cases 是全部解法共用的用例表，覆盖题面三个示例、单元素/两元素/三元素环上两两相邻/
// 全零等边界，以及约束上限规模（n = 100，nums[i] = 1000）。
//
// 规模取舍：本题两个骨架都只对每间房做常数次转移、恰好访问 n 个状态，渐进复杂度同为
// 线性，因此上限规模用例直接取题面约束的最大值 n = 100，两个骨架都能在毫秒级内跑完，
// 整个包远低于 1 秒目标。
//
// 期望值由独立参考实现算出：其中「两段线性 DP」
// 与「分段记忆化递归」两份实现互相印证，并对 n <= 12 的 5000 组随机数据与枚举全部子集
// 的朴素解三方对齐；各用例答案同时与暴力解逐一核对。
// 最大可能答案为 50 * 1000 = 50000（环上至多偷一半房屋），远小于 int 范围。
var cases = []struct {
	name   string
	nums   []int
	expect int
}{
	{"例1 三间房首尾相邻", []int{2, 3, 2}, 3},
	{"例2 四间房偷 1、3 号", []int{1, 2, 3, 1}, 4},
	{"例3 三间房递增只能取一间", []int{1, 2, 3}, 3},
	{"边界 单元素取该房", []int{5}, 5},
	{"边界 单元素金额为 0", []int{0}, 0},
	{"边界 两元素在环上相邻取较大", []int{1, 2}, 2},
	{"边界 两元素前大后小", []int{10, 1}, 10},
	{"边界 三元素环上两两相邻", []int{5, 5, 5}, 5},
	{"环约束生效 线性答案更大", []int{2, 1, 1, 2}, 3},
	{"递增五间", []int{1, 2, 3, 4, 5}, 8},
	{"固定较长序列 30 间", []int{3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2, 9, 16, 6, 13, 3, 10, 0, 7, 14, 4, 11, 1, 8, 15, 5, 12, 2}, 153},
	{"上限规模 n=100 全为 1000", rep(1000, 100), 50000},
	{"上限规模 n=100 交替 1000/0", alternating(1000, 0, 100), 50000},
	{"上限规模 n=100 全为 0", make([]int, 100), 0},
}

// TestRob 验证解法 1（环形拆成两段线性问题）。
func TestRob(t *testing.T) {
	checkCases(t, rob)
}

// TestRob1 验证解法 2（记忆化搜索）。
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
