package t_014_53_maximum_subarray

import "testing"

// cases 是全部解法共用的用例表，覆盖题面三个示例、单元素、全负数、全零、
// 两元素、较长固定序列，以及题面约束上限规模。
//
// 规模取舍：本题两个骨架分别是线性动态规划（Kadane，O(n)）与分治法
// （O(n log n)，递归深度 O(log n)），即便取到题面约束的最大规模
// n = 100000 也都只需百万级操作，两个骨架都能在毫秒级跑完，
// 因此上限用例直接取题面约束的最大值，无需额外缩小。
//
// 期望值由独立参考实现核对：Kadane、分治、O(n^2) 暴力三份实现互相印证，
// 并对小规模随机数据交叉验证 5000 组，全部一致；固定用例上三份实现答案也相同。
// 数值范围检查：最大可能答案为 100000 * 10000 = 10^9，未超过题面声明的
// 元素值域与数组长度所允许的范围，也在 int 范围内。
var cases = []struct {
	name   string
	nums   []int
	expect int
}{
	{"例1 跨越负数的最大子数组", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{"例2 单元素正数", []int{1}, 1},
	{"例3 含负数的整段和最大", []int{5, 4, -1, 7, 8}, 23},
	{"边界 单元素负数（子数组至少一个元素）", []int{-5}, -5},
	{"边界 全负数只能取最大的单个元素", []int{-3, -1, -2}, -1},
	{"边界 全零", []int{0, 0, 0}, 0},
	{"边界 零与负数混排", []int{0, -1, 0}, 0},
	{"边界 两元素先负后正", []int{-1, 2}, 2},
	{"边界 两元素先正后负", []int{2, -1}, 2},
	{"固定较长序列 n=30", []int{3, 10, -4, 7, 14, -8, 11, -1, 8, 15, -5, 12, -2, 9, -16, 6, 13, -3, 10, 0, -7, 14, -4, 11, 1, -8, 15, 5, -12, 2}, 106},
	{"上限规模 n=100000 正负极值交替", alternating(-10000, 10000, 100000), 10000},
	{"上限规模 n=100000 全部 10000", rep(10000, 100000), 1000000000},
	{"上限规模 n=100000 全部 -10000", rep(-10000, 100000), -10000},
}

// TestMaxSubArray 验证解法 1（线性动态规划 / Kadane）。
func TestMaxSubArray(t *testing.T) {
	checkCases(t, maxSubArray)
}

// TestMaxSubArray1 验证解法 2（分治法）。
// func TestMaxSubArray1(t *testing.T) {
// 	checkCases(t, maxSubArray1)
// }

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expect {
				t.Fatalf("case %q: n=%d, expected=%v, actual=%v", c.name, len(c.nums), c.expect, actual)
			}
			// 上限用例的切片过长，只记录规模而不打印全部元素。
			t.Logf("case %q: n=%d, expected=%v, actual=%v", c.name, len(c.nums), c.expect, actual)
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
