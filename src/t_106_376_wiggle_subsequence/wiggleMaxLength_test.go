package t_106_376_wiggle_subsequence

import "testing"

// wiggleCases 覆盖题面全部示例与各类边界情况。
// 题面约束为 1 <= nums.length <= 1000，故不包含空数组；
// wiggleMaxLength2 对空数组会因 up[0] 越界而 panic，两个实现均未做长度守卫。
var wiggleCases = []struct {
	name     string
	nums     []int
	expected int
}{
	{"例1 完整摆动序列", []int{1, 7, 4, 9, 2, 5}, 6},
	{"例2 含多个等长最优解", []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	{"例3 单调递增", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	{"单调递减", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 2},
	{"单元素", []int{5}, 1},
	{"两元素相等", []int{3, 3}, 1},
	{"两元素不等", []int{3, 5}, 2},
	{"全部相等", []int{7, 7, 7, 7}, 1},
	{"先平后摆", []int{3, 3, 3, 2, 5}, 3},
	{"平台后转折", []int{1, 2, 2, 2, 1}, 3},
	{"完全交替", []int{1, 2, 1, 2, 1, 2}, 6},
	{"长平台", []int{1, 1, 1, 2, 2, 2, 1, 1, 1}, 3},
}

// TestWiggleMaxLength 验证贪心实现，时间 O(n)、空间 O(1)。
func TestWiggleMaxLength(t *testing.T) {
	checkWiggleCases(t, wiggleMaxLength)
}

// TestWiggleMaxLength2 验证动态规划实现，时间 O(n)、空间 O(n)。
func TestWiggleMaxLength2(t *testing.T) {
	checkWiggleCases(t, wiggleMaxLength2)
}

func checkWiggleCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range wiggleCases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expected {
				t.Fatalf("nums=%v, expected=%v, actual=%v", c.nums, c.expected, actual)
			}
			t.Logf("nums=%v, expected=%v, actual=%v", c.nums, c.expected, actual)
		})
	}
}
