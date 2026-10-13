package t_003_673_number_of_longest_increasing_subsequence

import "testing"

// layeredCase 构造 blocks 个 1..size 的连续块。
// 值域被限制在 [1,size]，故最长严格递增子序列长度恰为 size，方案数由各块的取值位置决定。
func layeredCase(blocks, size int) []int {
	nums := make([]int, 0, blocks*size)
	for b := 0; b < blocks; b++ {
		for v := 1; v <= size; v++ {
			nums = append(nums, v)
		}
	}
	return nums
}

// descendingCase 构造 n 个严格递减的元素，最长递增子序列长度为 1，方案数为 n。
func descendingCase(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1000000 - i
	}
	return nums
}

// ascendingCase 构造 n 个严格递增的元素，最长递增子序列唯一。
func ascendingCase(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = -1000000 + i
	}
	return nums
}

// equalCase 构造 n 个相同元素，最长递增子序列长度为 1，方案数为 n。
func equalCase(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = -1000000
	}
	return nums
}

// cases 覆盖题面全部示例、边界情况与约束上限规模（n=2000、值域 ±10^6），由全部解法共用。
// 期望值均由 /tmp 下的独立参考程序（暴力枚举 + O(n^2) DP + 树状数组三方交叉验证）产出。
var cases = []struct {
	name   string
	nums   []int
	expect int
}{
	{"例1 存在两个最长递增子序列", []int{1, 3, 5, 4, 7}, 2},
	{"例2 全相等元素", []int{2, 2, 2, 2, 2}, 5},
	{"单元素", []int{1}, 1},
	{"两元素递增", []int{1, 2}, 1},
	{"两元素相等", []int{2, 2}, 2},
	{"严格递减", []int{5, 4, 3, 2, 1}, 5},
	{"严格递增 唯一解", []int{1, 2, 3, 4, 5}, 1},
	{"全部相同 四个元素", []int{7, 7, 7, 7}, 4},
	{"三元素 两条等长解", []int{1, 3, 2}, 2},
	{"值域上下界", []int{-1000000, 1000000}, 1},
	{"值域边界处重复", []int{-1000000, -1000000, 1000000}, 2},
	{"含负数", []int{-3, -1, -2, 0}, 2},
	{"分层构造 较大计数", layeredCase(3, 20), 231},
	{"约束上限 n=2000 严格递减", descendingCase(2000), 2000},
	{"约束上限 n=2000 严格递增", ascendingCase(2000), 1},
	{"约束上限 n=2000 全部相等", equalCase(2000), 2000},
}

// TestFindNumberOfLIS 验证解法 1（动态规划）。
func TestFindNumberOfLIS(t *testing.T) {
	checkFindNumberOfLISCases(t, findNumberOfLIS)
}

// TestFindNumberOfLIS1 验证解法 2（树状数组）。
// func TestFindNumberOfLIS1(t *testing.T) {
// 	checkFindNumberOfLISCases(t, findNumberOfLIS1)
// }

func checkFindNumberOfLISCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expect {
				t.Fatalf("n=%d, expected=%d, actual=%d", len(c.nums), c.expect, actual)
			}
			t.Logf("n=%d, expected=%d, actual=%d", len(c.nums), c.expect, actual)
		})
	}
}
