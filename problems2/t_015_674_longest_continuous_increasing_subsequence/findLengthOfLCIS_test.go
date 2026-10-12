package t_015_674_longest_continuous_increasing_subsequence

import (
	"fmt"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= nums.length <= 1e4，两个骨架（一次遍历、动态规划）的时间复杂度处于同一量级，
// 上限规模取到约束上限 1e4 时两者都能在 1 秒内轻松跑完，因此这里不再额外压低规模。

// incSeq 返回长度为 n 的严格递增数列，元素落在题面值域 [-10^9, 10^9] 内。
func incSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i - n/2
	}
	return s
}

// decSeq 返回长度为 n 的严格递减数列，元素落在题面值域 [-10^9, 10^9] 内。
func decSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = n/2 - i
	}
	return s
}

// eqSeq 返回长度为 n 的常数数列，用于覆盖「全部相等」的平台期边界。
func eqSeq(n int, v int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// zigSeq 返回长度为 n 的锯齿数列（严格交替升降），每段连续递增长度均为 2。
func zigSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = i / 2
		} else {
			s[i] = -i/2 - 1
		}
	}
	return s
}

// cases 覆盖题面全部示例与边界情况，由全部解法共用。
// 题面约束为 1 <= nums.length <= 1e4，故不含空数组；
// expected 均由 /tmp 下的独立参考实现（一次遍历 / 朴素枚举 / 动态规划三份交叉验证）跑出，
// 三份结果完全一致；数值范围检查通过：各用例元素均落在 [-10^9, 10^9] 内，
// 答案不超过数组长度 1e4，未超出 int 范围。
var cases = []struct {
	name     string
	nums     []int
	expected int
}{
	{"例1 递增段被 4 打断", []int{1, 3, 5, 4, 7}, 3},
	{"例2 全部相等", []int{2, 2, 2, 2, 2}, 1},
	{"单元素", []int{1}, 1},
	{"两元素递增", []int{1, 2}, 2},
	{"两元素递减", []int{2, 1}, 1},
	{"全程严格递增", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	{"全程严格递减", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1},
	{"值域两端极值相邻递增", []int{-1000000000, 1000000000}, 2},
	{"含平台期，递增段接在相等元素之后", []int{1, 2, 2, 3, 4}, 3},
	{"锯齿交替，每段长度均为 2", []int{1, 2, 1, 2, 1, 2}, 2},
	{"递增段出现在尾部之前", []int{-3, -1, 0, 2, -5}, 4},
	{"约束上限规模 1e4 严格递增", incSeq(10000), 10000},
	{"约束上限规模 1e4 严格递减", decSeq(10000), 1},
	{"约束上限规模 1e4 全部相等", eqSeq(10000, 7), 1},
	{"约束上限规模 1e4 锯齿交替", zigSeq(10000), 2},
}

// TestFindLengthOfLCIS 验证解法 1（一次遍历）。
func TestFindLengthOfLCIS(t *testing.T) {
	checkLCISCases(t, findLengthOfLCIS)
}

// TestFindLengthOfLCIS1 验证解法 2（动态规划）。
func TestFindLengthOfLCIS1(t *testing.T) {
	checkLCISCases(t, findLengthOfLCIS1)
}

func checkLCISCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expected {
				t.Fatalf("nums=%s, expected=%v, actual=%v", numsString(c.nums), c.expected, actual)
			}
			t.Logf("nums=%s, expected=%v, actual=%v", numsString(c.nums), c.expected, actual)
		})
	}
}

// numsString 在用例规模较大时只打印长度，避免失败信息过长。
func numsString(nums []int) string {
	if len(nums) > 20 {
		return fmt.Sprintf("<len=%d>", len(nums))
	}
	return fmt.Sprintf("%v", nums)
}
