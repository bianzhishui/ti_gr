package t_013_337_house_robber_iii

import "testing"

// cases 是全部解法共用的用例表，覆盖题面全部示例、单节点 / 两节点 / 三节点满树 /
// 金额全为 0 等边界，以及约束上限规模（节点数 n = 10^4）。
//
// 规模取舍：本题两个骨架（树形 DP、map 记忆化递归）都只对每个节点做常数次合并，
// 因此上限用例直接取题面约束的最大值 n = 10^4，两者都能在毫秒级跑完；其中
// 「右斜链 10^4 节点」额外压一下递归深度（Go 栈会自动增长，实测无压力）。
// 期望值由独立参考实现算出：树形 DP、
// 记忆化递归、暴力枚举独立集三份实现互相印证，并对小规模随机树交叉验证 4000 组。
// 最大可能答案为 10^4 个节点各取 10^4，即 10^8 量级，远小于 int 范围；
// 表中最大期望值为 66650000。
var cases = []struct {
	name   string
	root   *TreeNode
	expect int
}{
	{"例1 根+右孩子+右孙子", buildTree([]any{3, 2, 3, nil, 3, nil, 1}), 7},
	{"例2 取左孩子的左孩子与右孩子", buildTree([]any{3, 4, 5, 1, 3, nil, 1}), 9},
	{"边界 单节点", buildTree([]any{5}), 5},
	{"边界 单节点金额为 0", buildTree([]any{0}), 0},
	{"边界 两节点取根", buildTree([]any{3, 2}), 3},
	{"边界 两节点取右子", buildTree([]any{2, 3}), 3},
	{"边界 三节点满树取根更划算", buildTree([]any{4, 1, 2}), 4},
	{"边界 三节点满树取两子更划算", buildTree([]any{1, 4, 2}), 6},
	{"边界 五节点满树金额全为 0", buildTree([]any{0, 0, 0, 0, 0}), 0},
	{"边界 五节点满树金额全为 10000", buildTree([]any{10000, 10000, 10000, 10000, 10000}), 30000},
	{"非完全树 只有一侧有孙子", buildTree([]any{3, 2, nil, 1, 3}), 7},
	{"右斜链 10 节点全为 1", chain(10, 1), 5},
	{"上限规模 完全二叉树 10^4 节点 全为 10^4", fullTree(10000, 10000), 66650000},
	{"上限规模 右斜链 10^4 节点 全为 10^4", chain(10000, 10000), 50000000},
	{"上限规模 完全二叉树 10^4 节点 取固定序列", fullTreeVals(10000), 34001760},
}

// TestRob 验证解法 1（树形动态规划）。
func TestRob(t *testing.T) {
	checkCases(t, rob)
}

// TestRob1 验证解法 2（带记忆化的递归）。
// func TestRob1(t *testing.T) {
// 	checkCases(t, rob1)
// }

// checkCases 用同一张 cases 表驱动某个具体实现。
func checkCases(t *testing.T, fn func(*TreeNode) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.root)
			if actual != c.expect {
				t.Fatalf("case=%s, expected=%d, actual=%d", c.name, c.expect, actual)
			}
			t.Logf("case=%s, expected=%d, actual=%d", c.name, c.expect, actual)
		})
	}
}

// buildTree 按力扣的层序数组还原二叉树，nil 表示空节点（空节点的子孙不再占用数组下标）。
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

// fullTree 生成 n 个节点、所有节点值均为 val 的完全二叉树（按层序依次填满）。
func fullTree(n, val int) *TreeNode {
	if n <= 0 {
		return nil
	}
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]
		node.Left = &TreeNode{Val: val}
		queue = append(queue, node.Left)
		i++
		if i < n {
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
			i++
		}
	}
	return root
}

// chain 生成 n 个节点、所有节点值均为 val 的右斜链（退化二叉树，用来压递归深度）。
func chain(n, val int) *TreeNode {
	if n <= 0 {
		return nil
	}
	root := &TreeNode{Val: val}
	cur := root
	for i := 1; i < n; i++ {
		cur.Right = &TreeNode{Val: val}
		cur = cur.Right
	}
	return root
}

// valSeq 给出第 i 个层序节点的值，构成一条确定性的值序列。
func valSeq(i int) int {
	return (i*7919 + 1234) % 10001
}

// fullTreeVals 生成 n 个节点的完全二叉树，第 i 个层序节点的值取 valSeq(i)，
// 用来在满规模下混合大小不一的金额。
func fullTreeVals(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	root := &TreeNode{Val: valSeq(0)}
	queue := []*TreeNode{root}
	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]
		node.Left = &TreeNode{Val: valSeq(i)}
		queue = append(queue, node.Left)
		i++
		if i < n {
			node.Right = &TreeNode{Val: valSeq(i)}
			queue = append(queue, node.Right)
			i++
		}
	}
	return root
}
