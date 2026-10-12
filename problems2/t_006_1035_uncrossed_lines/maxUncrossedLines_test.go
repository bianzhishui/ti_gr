package t_006_1035_uncrossed_lines

import "testing"

// cases 覆盖题面全部示例与边界情况，并包含约束上限规模（两数组各 500 个元素），由全部解法共用。
// 两个骨架的时间复杂度同为 O(len(nums1)*len(nums2))，500x500 的用例在最坏情况下也只有 25 万次转移，
// 因此上限规模直接取题目声明的最值，无需为某个骨架缩小规模。
// 期望值由独立参考实现算出（二维 DP 与滚动数组两份实现互相印证），
// 全部答案均不超过 min(len(nums1), len(nums2)) <= 500，符合题目声明的数值范围。
var cases = []struct {
	name   string
	nums1  []int
	nums2  []int
	expect int
}{
	{"例1 交叉的两条线", []int{1, 4, 2}, []int{1, 2, 4}, 2},
	{"例2 含重复元素", []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
	{"例3 含重复元素与无公共元素段", []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	{"边界 单元素且相同", []int{1}, []int{1}, 1},
	{"边界 单元素且不同", []int{1}, []int{2}, 0},
	{"边界 无公共元素", []int{1, 2, 3}, []int{4, 5, 6}, 0},
	{"边界 全部相同但长度不等", []int{2, 2, 2}, []int{2, 2}, 2},
	{"边界 一个是另一个的子序列", []int{1, 2, 3}, []int{1, 2, 3, 4, 5}, 3},
	{"边界 相同元素重复出现", []int{5, 5, 5, 5}, []int{5, 5, 5}, 3},
	{"经典交错序列", []int{1, 3, 4, 9, 5}, []int{1, 4, 9, 5, 3}, 4},
	{"上限规模 500x500 全相同", filled(500, 7), filled(500, 7), 500},
	{"上限规模 500x500 完全不同", filled(500, 7), filled(500, 8), 0},
	{"上限规模 500x500 周期串", repeatPattern([]int{1, 2}, 250), repeatPattern([]int{2, 1}, 250), 499},
}

// TestMaxUncrossedLines 验证解法 1（二维动态规划）。
func TestMaxUncrossedLines(t *testing.T) {
	checkCases(t, maxUncrossedLines)
}

// TestMaxUncrossedLines1 验证解法 2（滚动数组）。
func TestMaxUncrossedLines1(t *testing.T) {
	checkCases(t, maxUncrossedLines1)
}

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int, []int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums1, c.nums2)
			if actual != c.expect {
				t.Fatalf("nums1=%v, nums2=%v, expected=%v, actual=%v", c.nums1, c.nums2, c.expect, actual)
			}
			t.Logf("nums1=%v, nums2=%v, expected=%v, actual=%v", c.nums1, c.nums2, c.expect, actual)
		})
	}
}

// filled 返回长度为 n、每个元素都是 v 的切片，用于构造上限规模用例。
func filled(n int, v int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}

// repeatPattern 返回把 pattern 重复 times 次拼接得到的切片，用于构造上限规模用例。
func repeatPattern(pattern []int, times int) []int {
	out := make([]int, 0, len(pattern)*times)
	for i := 0; i < times; i++ {
		out = append(out, pattern...)
	}
	return out
}
