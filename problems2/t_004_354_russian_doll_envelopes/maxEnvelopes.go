package t_004_354_russian_doll_envelopes

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

// 解法 1：排序（宽度升序、同宽度高度降序）后对高度求最长严格递增子序列（贪心 + 二分 O(n log n)）
// TODO: 实现
func maxEnvelopes(envelopes [][]int) int {
	// 1. 排序：宽升序；宽相同则高降序
	slices.SortFunc(envelopes, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(a[0], b[0]), // 宽升序
			cmp.Compare(b[1], a[1]), // 同宽，高降序
		)
	})

	// 2. 对高度数组求最长严格递增子序列（LIS）
	// 这里用 O(n log n) 的耐心排序，也可以用 O(n²) 的 DP
	var tails []int // tails[k] = 长度为 k+1 的递增子序列的最小结尾高度
	for _, env := range envelopes {
		h := env[1]
		// 二分找到第一个 >= h 的位置
		i := sort.SearchInts(tails, h)
		if i == len(tails) {
			tails = append(tails, h) // 接在末尾，长度+1
		} else {
			tails[i] = h // 替换，让结尾更小
		}
	}
	return len(tails)
}

// 解法 2：排序后对高度做 O(n^2) 的动态规划
// TODO: 实现
func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	// 宽降序；宽相同则高升序
	slices.SortFunc(envelopes, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(b[0], a[0]), // 宽降序
			cmp.Compare(a[1], b[1]), // 高升序
		)
	})

	for i, e := range envelopes {
		fmt.Printf("envelopes[%d] = [%d, %d]\n", i, e[0], e[1])
	}

	n := len(envelopes)
	dp := make([]int, n)
	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 宽降序下：j 的宽 >= i 的宽
			// 要 i 能套住 j，需 i 更小：宽严格小于，高严格小于
			if envelopes[j][0] > envelopes[i][0] &&
				envelopes[j][1] > envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

func maxEnvelopes2(envelopes [][]int) int {
	slices.SortFunc(envelopes, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(a[0], b[0]), // 宽升序
			cmp.Compare(b[1], a[1]), // 同宽，高降序
		)
	})

	var tails []int
	for _, env := range envelopes {
		h := env[1]

		// 用 slices.BinarySearch 找插入/替换位置
		i, found := slices.BinarySearch(tails, h)
		if found {
			// 找到了相等的 h，位置 i 就是第一个 >= h 的位置，直接替换
			tails[i] = h
		} else if i == len(tails) {
			// 没找到且 i 在末尾，说明 h 比所有元素都大，追加
			tails = append(tails, h)
		} else {
			// 没找到且 i 在中间，i 是第一个 > h 的位置，替换它
			tails[i] = h
		}
	}
	return len(tails)
}
