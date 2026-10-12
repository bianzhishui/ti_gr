package t_023_446_arithmetic_slices_ii_subsequence

import "testing"

// 用例表由全部解法共用。
// 规模取舍：本题两种骨架（哈希表 DP、按 (下标, 公差) 记忆化搜索）的可行实现均为 O(n^2)
// 级别，因此上限规模用例直接取题面上限 n = 1000，任一骨架都能从容跑完（整个包 1 秒内）。
// 该 1000 长度的用例特意选取不会产生海量等差子序列的数列，以遵守题面
// 「答案是一个 32-bit 整数」的保证，同时避免答案溢出。
// 全部期望值均由独立参考实现核对得到（两套不同思路的参考实现结果一致），未凭记忆填写。

type testCase struct {
	name string
	nums []int
	want int
}

// genUpperLimit 生成 n = 1000 的上限规模用例：nums[i] = (i^3 + 7i + 3) mod 100003。
func genUpperLimit() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i*i*i + 7*i + 3) % 100003
	}
	return nums
}

var cases = []testCase{
	{"示例 1：公差为 2 的等差数列", []int{2, 4, 6, 8, 10}, 7},
	{"示例 2：五个元素全部相等", []int{7, 7, 7, 7, 7}, 16},
	{"边界：只有一个元素", []int{1}, 0},
	{"边界：只有两个元素", []int{1, 2}, 0},
	{"边界：三个元素但不是等差", []int{1, 2, 4}, 0},
	{"边界：最短的等差子序列", []int{1, 2, 3}, 1},
	{"边界：四个元素全部相等", []int{0, 0, 0, 0}, 5},
	{"边界：包含负数", []int{-1, 1, 3, 5}, 3},
	{"边界：元素取到 int32 上下界", []int{-2147483648, 2147483647, -2147483648}, 0},
	{"上限规模：n = 1000", genUpperLimit(), 823},
}

// checkCases 用同一张用例表驱动任意一个解法实现。
func checkCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fn(tc.nums)
			if actual != tc.want {
				t.Fatalf("numberOfArithmeticSlices(n=%d) = %d, want %d", len(tc.nums), actual, tc.want)
			}
			t.Logf("numberOfArithmeticSlices(n=%d) = %d", len(tc.nums), actual)
		})
	}
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	checkCases(t, numberOfArithmeticSlices)
}

func TestNumberOfArithmeticSlices1(t *testing.T) {
	checkCases(t, numberOfArithmeticSlices1)
}
