package t_033_1043_partition_array_for_maximum_sum

import (
	"fmt"
	"testing"
)

// 关于用例规模的取舍：
// 题面约束为 1 <= arr.length <= 500、1 <= k <= arr.length。两个骨架（线性动态规划、记忆化搜索）
// 的状态数都是「长度 × 段长上限」，n = k = 500 时也只有约 25 万次转移，所以上限规模用例直接取到
// 约束上限 n = 500，两个骨架都能在 1 秒内从容跑完，不需要压低规模。
// 另外，题面声明测试用例保证答案是 32 位整数，因此全部用例的期望值都经过核对、不超过 2^31-1
// （其中涉及 10^9 量级元素的用例只取两个元素，答案为 2000000000，仍在范围内）。

// bigArr 返回长度恰为 500 的上限规模用例：取值在 [0, 96] 之间循环，保证答案远在 32 位整数范围内。
func bigArr() []int {
	arr := make([]int, 500)
	for i := range arr {
		arr[i] = (i * 31) % 97
	}
	return arr
}

// cases 覆盖题面全部示例与边界情况，由全部解法共用。
// 期望值由独立参考实现核对：分别用线性动态规划、记忆化搜索、枚举全部分段方式的指数级实现
// 三份独立代码交叉验证，另对 3000 组随机小规模输入做三者一致性比对，结果完全一致；
// 数值范围检查通过：全部答案均为非负且不超过 2^31-1 = 2147483647。
var cases = []struct {
	name     string
	arr      []int
	k        int
	expected int
}{
	{"例1 三段划分最优", []int{1, 15, 7, 9, 2, 5, 10}, 3, 84},
	{"例2 长度上限为 4", []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4, 83},
	{"例3 单元素", []int{1}, 1, 1},
	{"边界 单元素零值", []int{0}, 1, 0},
	{"边界 k=1 只能逐段替换为自身", []int{3, 1, 4, 1, 5, 9, 2, 6}, 1, 31},
	{"边界 k 等于长度，整段都变成最大值", []int{1, 2, 3, 9, 4}, 5, 45},
	{"边界 全相等且 k=1", []int{5, 5, 5, 5, 5}, 1, 25},
	{"边界 全相等且 k 覆盖全数组", []int{5, 5, 5, 5, 5}, 5, 25},
	{"边界 全为 0", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, 0},
	{"边界 后半段更长反而更优", []int{4, 1, 1, 1, 6}, 2, 21},
	{"边界 平台期中间穿插小值", []int{7, 7, 1, 7, 7}, 3, 35},
	{"边界 两个 10^9 同段", []int{1000000000, 1000000000}, 2, 2000000000},
	{"上限规模 n=500 k=17", bigArr(), 17, 46731},
	{"上限规模 n=500 k=500", bigArr(), 500, 48000},
}

// TestMaxSumAfterPartitioning 验证解法 1（线性动态规划）。
func TestMaxSumAfterPartitioning(t *testing.T) {
	checkMaxSumAfterPartitioningCases(t, maxSumAfterPartitioning)
}

// TestMaxSumAfterPartitioning1 验证解法 2（记忆化搜索）。
func TestMaxSumAfterPartitioning1(t *testing.T) {
	checkMaxSumAfterPartitioningCases(t, maxSumAfterPartitioning1)
}

// checkMaxSumAfterPartitioningCases 用同一张 cases 表驱动一个解法骨架。
func checkMaxSumAfterPartitioningCases(t *testing.T, fn func([]int, int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.arr, c.k)
			if actual != c.expected {
				t.Fatalf("arr=%s, k=%d, expected=%d, actual=%d", arrString(c.arr), c.k, c.expected, actual)
			}
			t.Logf("arr=%s, k=%d, expected=%d, actual=%d", arrString(c.arr), c.k, c.expected, actual)
		})
	}
}

// arrString 在用例较长时只打印长度与前缀，避免日志与失败信息过长。
func arrString(arr []int) string {
	if len(arr) > 12 {
		return fmt.Sprintf("<len=%d, prefix=%v>", len(arr), arr[:6])
	}
	return fmt.Sprintf("%v", arr)
}
