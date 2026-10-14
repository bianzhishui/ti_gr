package main

import (
	"reflect"
	"testing"
)

// rep 返回长度为 n、元素全为 v 的切片，用于构造约束上限规模的用例。
func rep(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// cases 覆盖题面示例、边界情况与约束上限规模，由全部解法共用。
// 期望值由独立参考实现核对（暴力枚举、空隙中点、大段模拟三份实现交叉验证）。
// 题面约束为 1 <= seatNum <= 500，用例规模直接取到上限，三个解法都能在毫秒级跑完。
//
// 语义说明（题面未完全明确之处，按此实现）：座位的输出取「最后一次进场动作」的结果，
// 若该次进场时座位已满则输出 -1；中途某次进场满员、但之后有人离场又有新员工成功落座时，
// 输出的是后者的座位号（见用例「中途满员但最后一次进场成功」）。
var cases = []struct {
	name     string
	seatNum  int
	seq      []int
	expected int
}{
	{"示例1 进进进进离进", 10, []int{1, 1, 1, 1, -4, 1}, 5},
	{"单座位进1次", 1, []int{1}, 0},
	{"单座位坐满后再进", 1, []int{1, 1}, -1},
	{"5座依次坐满", 5, []int{1, 1, 1, 1, 1}, 3},
	{"5座坐满后再进", 5, []int{1, 1, 1, 1, 1, 1}, -1},
	{"5座满员后离场再进", 5, []int{1, 1, 1, 1, 1, -1, 1}, 1},
	{"中途满员但最后一次进场成功", 2, []int{1, 1, 1, -1, 1}, 1},
	{"3座依次进满", 3, []int{1, 1, 1}, 1},
	{"7座进4人", 7, []int{1, 1, 1, 1}, 1},
	{"10座进3人并列取最小索引", 10, []int{1, 1, 1}, 4},
	{"右端离场后重新可用", 10, []int{1, 1, -9, 1}, 9},
	{"示例1再增两步", 10, []int{1, 1, 1, 1, -4, 1, -9, 1}, 9},
	{"上限规模 500座进499人", 500, rep(1, 499), 496},
	{"上限规模 500座进500人", 500, rep(1, 500), 498},
	{"上限规模 500座进501人", 500, rep(1, 501), -1},
}

// TestMaxSocialDistance 验证解法 1（暴力枚举每个空座位）。
func TestMaxSocialDistance(t *testing.T) {
	checkCases(t, maxSocialDistance)
}

// TestMaxSocialDistance1 验证解法 2（有序已坐座位 + 空隙中点）。
func TestMaxSocialDistance1(t *testing.T) {
	checkCases(t, maxSocialDistance1)
}

// TestMaxSocialDistance2 验证解法 3（大顶堆维护空闲区间）。
func TestMaxSocialDistance2(t *testing.T) {
	checkCases(t, maxSocialDistance2)
}

func checkCases(t *testing.T, fn func(int, []int) int) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fn(c.seatNum, c.seq)
			if actual != c.expected {
				t.Fatalf("seatNum=%d, len(seq)=%d, expected=%d, actual=%d", c.seatNum, len(c.seq), c.expected, actual)
			}
			t.Logf("seatNum=%d, len(seq)=%d, expected=%d, actual=%d", c.seatNum, len(c.seq), c.expected, actual)
		})
	}
}

// TestParseInput 验证输入解析：兼容中括号、逗号、空白分隔等写法。
func TestParseInput(t *testing.T) {
	parseCases := []struct {
		name     string
		input    string
		wantSeat int
		wantSeq  []int
		wantErr  bool
	}{
		{"题面格式", "10\n[1,1,1,1,-4,1]", 10, []int{1, 1, 1, 1, -4, 1}, false},
		{"空格分隔数组", "10\n1 1 1 1 -4 1", 10, []int{1, 1, 1, 1, -4, 1}, false},
		{"逗号后带空格", "10\n[1, 1, 1]", 10, []int{1, 1, 1}, false},
		{"单元素数组", "1\n[1]", 1, []int{1}, false},
		{"只有座位数", "5", 5, []int{}, false},
		{"结尾换行", "10\n[1,1,1,1,-4,1]\n", 10, []int{1, 1, 1, 1, -4, 1}, false},
		{"空输入", "", 0, nil, true},
		{"非数字", "abc", 0, nil, true},
	}

	for _, c := range parseCases {
		t.Run(c.name, func(t *testing.T) {
			seatNum, seq, err := parseInput(c.input)
			if c.wantErr {
				if err == nil {
					t.Fatalf("input=%q, 期望报错但成功返回", c.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("input=%q, 意外报错: %v", c.input, err)
			}
			if seatNum != c.wantSeat {
				t.Fatalf("input=%q, seatNum 期望 %d, 实际 %d", c.input, c.wantSeat, seatNum)
			}
			if !reflect.DeepEqual(seq, c.wantSeq) {
				t.Fatalf("input=%q, seq 期望 %v, 实际 %v", c.input, c.wantSeq, seq)
			}
			t.Logf("input=%q -> seatNum=%d, seq=%v", c.input, seatNum, seq)
		})
	}
}
