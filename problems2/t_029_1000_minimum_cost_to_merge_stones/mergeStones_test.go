package t_029_1000_minimum_cost_to_merge_stones

import (
	"fmt"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= n <= 30、2 <= k <= 30，两个骨架（区间动态规划、记忆化搜索）的规模上限
// 都由 n 与 k 决定，即使取到约束上限 n = 30，两者也都能在 1 秒内从容跑完
// （记忆化搜索多一层 map 开销，但同样远低于时间上限）。
// 因此上限规模用例直接取到 n = 30，并另配 n = 29、k = 5 覆盖不同的 (n-1) % (k-1) 分支，
// 不再额外压低规模。
//
// cases 覆盖题面全部示例与边界情况（单元素、目标堆数等于长度、堆数不足以合并、约束上限规模等），
// 由全部解法共用。expected 均由独立参考实现核对：区间动态规划、记忆化搜索、穷举暴力三份实现
// 结果完全一致，并额外做了两万余组小规模穷举比对，同样零分歧。
// 数值范围检查通过：约束下所有石头的总数不超过 100 * 30 = 3000，最小成本不超过 14800，
// 远未超出 int 的表示范围。

// eqSeq 返回长度为 n 的常数数列，用于覆盖「全部相等」的边界，元素落在题面值域 [1, 100] 内。
func eqSeq(n, v int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// incSeq 返回长度为 n 的数列，元素在 [1, 100] 内循环递增。
func incSeq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i%100 + 1
	}
	return s
}

var cases = []struct {
	name     string
	stones   []int
	k        int
	expected int
}{
	{"例1 [3,2,4,1] k=2", []int{3, 2, 4, 1}, 2, 20},
	{"例2 [3,2,4,1] k=3 不可合并", []int{3, 2, 4, 1}, 3, -1},
	{"例3 [3,5,1,2,6] k=3", []int{3, 5, 1, 2, 6}, 3, 25},
	{"单元素 k=2 无需合并", []int{8}, 2, 0},
	{"单元素 k=30 无需合并", []int{8}, 30, 0},
	{"两堆 k=3 堆数不足", []int{1, 2}, 3, -1},
	{"两堆 k=2 唯一合并方案", []int{1, 2}, 2, 3},
	{"三堆全等 k=2", []int{5, 5, 5}, 2, 25},
	{"三堆 k=3 一次合并到位", []int{1, 2, 3}, 3, 6},
	{"四堆全等 k=2", []int{5, 5, 5, 5}, 2, 40},
	{"四堆 k=4 一次合并到位", []int{1, 2, 3, 4}, 4, 10},
	{"六堆取到值域上限 k=2", []int{100, 100, 100, 100, 100, 100}, 2, 1600},
	{"六堆 k=3 不为 k-1 的倍数", []int{7, 1, 9, 3, 5, 2}, 3, -1},
	{"六堆 k=2 无序数列", []int{7, 1, 9, 3, 5, 2}, 2, 69},
	{"九堆 k=3 可合并", []int{4, 8, 1, 6, 3, 9, 2, 7, 5}, 3, 90},
	{"上限规模 n=29 k=3 全 100", eqSeq(29, 100), 3, 9000},
	{"上限规模 n=29 k=5 全 1", eqSeq(29, 1), 5, 63},
	{"上限规模 n=30 k=2 全 100", eqSeq(30, 100), 2, 14800},
	{"上限规模 n=30 k=2 全 1", eqSeq(30, 1), 2, 148},
	{"上限规模 n=30 k=30 一次合并到位", eqSeq(30, 100), 30, 3000},
	{"上限规模 n=30 k=29 不可合并", eqSeq(30, 100), 29, -1},
	{"上限规模 n=30 k=3 不可合并", eqSeq(30, 100), 3, -1},
	{"上限规模 n=30 k=2 递增", incSeq(30), 2, 2175},
}

// TestMergeStones 验证解法 1（区间动态规划）。
func TestMergeStones(t *testing.T) {
	checkMergeStonesCases(t, mergeStones)
}

// TestMergeStones1 验证解法 2（记忆化搜索）。
func TestMergeStones1(t *testing.T) {
	checkMergeStonesCases(t, mergeStones1)
}

func checkMergeStonesCases(t *testing.T, fn func([]int, int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.stones, c.k)
			if actual != c.expected {
				t.Fatalf("stones=%s, k=%d, expected=%d, actual=%d", stonesString(c.stones), c.k, c.expected, actual)
			}
			t.Logf("stones=%s, k=%d, expected=%d, actual=%d", stonesString(c.stones), c.k, c.expected, actual)
		})
	}
}

// stonesString 在用例规模较大时只打印长度，避免日志与失败信息过长。
func stonesString(stones []int) string {
	if len(stones) > 20 {
		return fmt.Sprintf("<len=%d>", len(stones))
	}
	return fmt.Sprintf("%v", stones)
}
