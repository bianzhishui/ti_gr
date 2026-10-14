package t_016_718_maximum_length_of_repeated_subarray

import (
	"fmt"
	"testing"
)

// 两个骨架都是 O(len(nums1)*len(nums2)) 的填表做法，量级相同，
// 因此上限规模用例直接取题面允许的最大长度 1000（约 10^6 次状态转移），
// 保证整张表对每个骨架都能在 1 秒内从容跑完。
//
// cases 覆盖题面全部示例与边界情况（单元素、无公共元素、最长匹配不在开头、
// 周期串错位、约束上限规模），由全部解法共用。
// 期望值由独立参考实现核对（二维 DP、滚动数组、暴力枚举三份实现结果一致，
// 并另有随机小规模交叉验证），未凭记忆填写；所有答值均不超过 1000 的声明范围。
var cases = []struct {
	name   string
	nums1  []int
	nums2  []int
	expect int
}{
	{"例1 公共子数组 [3,2,1]", []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
	{"例2 五个 0 完全相同", []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	{"边界 单元素相同", []int{1}, []int{1}, 1},
	{"边界 单元素不同", []int{1}, []int{2}, 0},
	{"边界 单元素与长数组共享一个元素", []int{1}, []int{3, 1, 4}, 1},
	{"边界 无公共元素", []int{1, 2, 3}, []int{4, 5, 6}, 0},
	{"边界 全部元素都不匹配", []int{1, 1, 1, 1}, []int{2, 2, 2, 2}, 0},
	{"最长匹配不在开头", []int{5, 1, 2, 3, 4}, []int{0, 1, 2, 3, 9}, 3},
	{"交错重复 1..3", []int{1, 2, 3, 1, 2, 3, 4}, []int{2, 3, 1, 2, 3, 4}, 6},
	{"周期串 0101 与 1010 错位", cycle([]int{0, 1}, 8), cycle([]int{1, 0}, 8), 7},
	{"上限规模 1000x1000 全相同", repeatValue(7, 1000), repeatValue(7, 1000), 1000},
	{"上限规模 1000x1000 完全不同", repeatValue(0, 1000), repeatValue(1, 1000), 0},
	{"上限规模 1000x1000 周期串错位", cycle([]int{0, 1}, 1000), cycle([]int{1, 0}, 1000), 999},
	{"上限规模 1000x1000 仅尾部不同", repeatValue(0, 1000), append(repeatValue(0, 999), 1), 999},
}

// repeatValue 返回长度为 n、每个元素都是 v 的切片。
func repeatValue(v int, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// cycle 返回长度为 n、按 vals 循环填充的切片。
func cycle(vals []int, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = vals[i%len(vals)]
	}
	return s
}

// brief 把超长切片压缩成简短形式，避免失败信息打印上千个元素。
func brief(s []int) string {
	if len(s) <= 12 {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("len=%d %v...", len(s), s[:12])
}

// TestFindLength 验证解法 1（二维动态规划）。
func TestFindLength(t *testing.T) {
	checkCases(t, findLength)
}

// TestFindLength1 验证解法 2（滚动数组）。
// func TestFindLength1(t *testing.T) {
// 	checkCases(t, findLength1)
// }

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func([]int, []int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums1, c.nums2)
			if actual != c.expect {
				t.Fatalf("nums1=%s, nums2=%s, expected=%v, actual=%v", brief(c.nums1), brief(c.nums2), c.expect, actual)
			}
			t.Logf("expected=%v, actual=%v", c.expect, actual)
		})
	}
}
