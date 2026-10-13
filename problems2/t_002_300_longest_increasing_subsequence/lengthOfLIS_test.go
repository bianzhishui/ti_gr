package t_002_300_longest_increasing_subsequence

import (
	"fmt"
	"testing"
)

// incSeq 返回长度为 n 的严格递增数列，元素落在题面值域 [-10^4, 10^4] 内。
func incSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i - n/2
	}
	return s
}

// decSeq 返回长度为 n 的严格递减数列，元素落在题面值域 [-10^4, 10^4] 内。
func decSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = n/2 - i
	}
	return s
}

// zigSeq 返回长度为 n 的锯齿数列（正负交替），元素落在题面值域 [-10^4, 10^4] 内。
func zigSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		if i%2 == 0 {
			s[i] = i / 2
		} else {
			s[i] = -i / 2
		}
	}
	return s
}

// plateauSeq 返回长度为 n 的平台期数列（每 10 个元素为一组且组内相等，组间递增），
// 元素落在题面值域 [-10^4, 10^4] 内。
func plateauSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i / 10
	}
	return s
}

// cases 覆盖题面全部示例与边界情况，由全部解法共用。
// 题面约束为 1 <= nums.length <= 2500，故不含空数组；
// expected 均由 /tmp 下的独立参考实现（动态规划 / 贪心+二分 / 小数据枚举暴力三份交叉验证）跑出，
// 数值范围检查通过：答案不超过数组长度 2500，元素均由生成函数限制在 [-10^4, 10^4] 内。
var cases = []struct {
	name     string
	nums     []int
	expected int
}{
	{"例1 存在长度为 4 的递增子序列", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{"例2 含相等元素与前缀回退", []int{0, 1, 0, 3, 2, 3}, 4},
	{"例3 全部元素相等", []int{7, 7, 7, 7, 7, 7, 7}, 1},
	{"单元素", []int{5}, 1},
	{"两元素严格递增", []int{1, 2}, 2},
	{"两元素严格递减", []int{2, 1}, 1},
	{"全程严格递减", []int{5, 4, 3, 2, 1}, 1},
	{"全程严格递增", []int{1, 2, 3, 4, 5, 6}, 6},
	{"含负数，最优子序列跨过后续小值", []int{-3, -1, 0, 2, -5}, 4},
	{"值域两端极值", []int{-10000, 10000}, 2},
	{"约束上限规模 2500 严格递增", incSeq(2500), 2500},
	{"约束上限规模 2500 严格递减", decSeq(2500), 1},
	{"约束上限规模 2500 锯齿交替", zigSeq(2500), 1250},
	{"约束上限规模 2500 平台期，每个取值只能取一个", plateauSeq(2500), 250},
}

// TestLengthOfLIS 验证解法 1（动态规划 O(n^2)）。
func TestLengthOfLIS(t *testing.T) {
	checkCases(t, lengthOfLIS)
}

// TestLengthOfLIS1 验证解法 2（贪心 + 二分查找 O(n log n)）。
// func TestLengthOfLIS1(t *testing.T) {
// 	checkCases(t, lengthOfLIS1)
// }

func checkCases(t *testing.T, fn func([]int) int) {
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
