package t_001_377_combination_sum_iv

import "testing"

// cases 覆盖题面示例与边界情况，由全部解法共用。
// 新增用例时请确认所有实现都能在合理时间内完成。
var cases = []struct {
	name   string
	nums   []int
	target int
	expect int
}{
	{"例1 顺序不同视为不同组合", []int{1, 2, 3}, 4, 7},
	{"例2 无法凑出", []int{9}, 3, 0},
	{"单元素恰好命中", []int{1}, 1, 1},
	{"单元素需要取两次", []int{1}, 2, 1},
	{"目标小于所有元素", []int{3, 4}, 2, 0},
	{"同一元素取多次", []int{2}, 6, 1},
	{"两元素组合", []int{1, 2}, 3, 3},
	{"nums 顺序不影响结果", []int{2, 1, 3}, 4, 7},
	{"含间隔元素", []int{1, 2, 5}, 11, 218},
	{"较大目标", []int{1, 2, 3}, 20, 121415},
	{"约束上限 target=1000", []int{500, 1000}, 1000, 2},
}

// TestCombinationSum4 验证解法 1（记忆化搜索）。
func TestCombinationSum4(t *testing.T) {
	checkCases(t, combinationSum4)
}

// TestCombinationSum4_1 验证解法 2（完全背包 DP）。
func TestCombinationSum4_1(t *testing.T) {
	checkCases(t, combinationSum4_1)
}

func checkCases(t *testing.T, fn func([]int, int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.nums, c.target)
			if actual != c.expect {
				t.Fatalf("nums=%v, target=%v, expected=%v, actual=%v", c.nums, c.target, c.expect, actual)
			}
			t.Logf("nums=%v, target=%v, expected=%v, actual=%v", c.nums, c.target, c.expect, actual)
		})
	}
}
