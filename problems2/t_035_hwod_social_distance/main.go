// 社交距离（华为OD统一考试 C卷）——标准输入输出型题目。
//
// 输入：第一行为座位总数 seatNum，第二行为进出顺序数组 seatOrLeave，形如
//
//	10
//	[1,1,1,1,-4,1]
//
// 输出：最后进场员工落座的座位号；若他进场时座位已满，输出 -1。
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 解法 1：暴力枚举每个空座位
// 每次进场时遍历所有空位，用「到左侧最近已坐位的距离」与「到右侧最近已坐位的距离」的较小值
// 作为该座位的社交距离，取社交距离最大者，并列时取索引最小者。
// 房间全空时直接坐 0（题面规定的特殊情形）。整体 O(len(seatOrLeave) * seatNum)。
// TODO: 实现
func maxSocialDistance(seatNum int, seatOrLeave []int) int {
	return 0
}

// 解法 2：维护有序的已坐座位，只考察空隙中点
// 座位 0 一旦坐人就永不离开，因此左端不会空出；进场者的最优位置只可能是
// 某个「相邻已坐位之间空隙」的中点（空隙长度为偶数时两个中点并列，取索引较小者）
// 或最右端空位。每次进场只需比较这些候选，整体 O(len(seatOrLeave) * seatNum)。
// TODO: 实现
func maxSocialDistance1(seatNum int, seatOrLeave []int) int {
	return 0
}

// 解法 3：大顶堆维护空闲区间
// 把每个空闲区间按「最大社交距离」为优先级放进大顶堆（并列时取区间左端更小者），
// 进场时取出堆顶区间落座并把剩余部分重新入堆，离场时与左右相邻区间合并。
// 平均 O(len(seatOrLeave) * log seatNum)。
// TODO: 实现
func maxSocialDistance2(seatNum int, seatOrLeave []int) int {
	return 0
}

// parseInput 解析标准输入：第一个整数是座位总数，其余整数是进出顺序数组。
// 同时兼容 [1,1,1,1,-4,1]、1 1 1 1 -4 1、[1, 1, 1] 等写法。
func parseInput(input string) (seatNum int, seatOrLeave []int, err error) {
	s := strings.NewReplacer("[", " ", "]", " ", ",", " ", "\n", " ", "\r", " ", "\t", " ").Replace(input)
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return 0, nil, errors.New("输入为空")
	}

	nums := make([]int, 0, len(fields))
	for _, f := range fields {
		v, convErr := strconv.Atoi(f)
		if convErr != nil {
			return 0, nil, fmt.Errorf("无法解析 %q: %w", f, convErr)
		}
		nums = append(nums, v)
	}
	return nums[0], nums[1:], nil
}

func main() {
	data, readErr := io.ReadAll(os.Stdin)
	if readErr != nil {
		fmt.Fprintln(os.Stderr, "读取输入失败:", readErr)
		os.Exit(1)
	}

	seatNum, seatOrLeave, parseErr := parseInput(string(data))
	if parseErr != nil {
		fmt.Fprintln(os.Stderr, "解析输入失败:", parseErr)
		os.Exit(1)
	}

	fmt.Println(maxSocialDistance(seatNum, seatOrLeave))
}
