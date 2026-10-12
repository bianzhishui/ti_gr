package t_028_312_burst_balloons

import "testing"

// repeatInt 构造 n 个值为 v 的切片，供上限规模用例使用。
func repeatInt(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// maxCoinsCases 覆盖题面全部示例、单元素/全零/含零等边界情况，以及题面约束的上限规模。
// 两个骨架（区间动态规划、记忆化搜索）的渐进复杂度同阶，上限规模用例直接取题面允许的
// 最大值 n = 300，两者都能在整体 1 秒内从容跑完，因此不需要为某个骨架单独缩小规模。
// 全部期望值由独立参考实现核对（区间 DP 与记忆化搜索两种写法结果完全一致），
// 其中最大期望值 298010100 未超出题目声明的 32 位整数范围。
var maxCoinsCases = []struct {
	name     string
	nums     []int
	expected int
}{
	{"例1 四只气球", []int{3, 1, 5, 8}, 167},
	{"例2 两只气球", []int{1, 5}, 10},
	{"单元素 7", []int{7}, 7},
	{"单元素 0", []int{0}, 0},
	{"单元素 100", []int{100}, 100},
	{"三元素全零", []int{0, 0, 0}, 0},
	{"四元素全一", []int{1, 1, 1, 1}, 4},
	{"两只气球均为最大值", []int{100, 100}, 10100},
	{"严格递增", []int{1, 2, 3, 4, 5}, 110},
	{"含零元素", []int{0, 5, 0, 3}, 20},
	{"递增 20 个元素", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 45640},
	{"上限规模 n=300 全为 100", repeatInt(100, 300), 298010100},
}

// TestMaxCoins 验证区间动态规划骨架。
func TestMaxCoins(t *testing.T) {
	checkMaxCoinsCases(t, maxCoins)
}

// TestMaxCoins1 验证记忆化搜索骨架。
func TestMaxCoins1(t *testing.T) {
	checkMaxCoinsCases(t, maxCoins1)
}

// checkMaxCoinsCases 用同一张用例表驱动一个解法。
func checkMaxCoinsCases(t *testing.T, fn func([]int) int) {
	t.Helper()
	for _, c := range maxCoinsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums)
			if actual != c.expected {
				t.Fatalf("nums(n=%d)=%v, expected=%d, actual=%d", len(c.nums), c.nums, c.expected, actual)
			}
			t.Logf("n=%d, expected=%d, actual=%d", len(c.nums), c.expected, actual)
		})
	}
}
