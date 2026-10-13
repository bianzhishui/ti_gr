package t_006_1035_uncrossed_lines

import "slices"

// 解法 1：二维动态规划 —— 不相交的连线等价于在两数组中按顺序配对相等元素，即求两个数组的最长公共子序列长度
// TODO: 实现
func maxUncrossedLines(nums1 []int, nums2 []int) int {

	// nums1 = append([]int{0}, nums1...)
	// nums2 = append([]int{0}, nums2...)

	nums1 = slices.Insert(nums1, 0, 0)
	nums2 = slices.Insert(nums2, 0, 0)

	nums1len := len(nums1)
	nums2len := len(nums2)

	dp := make([][]int, nums1len)
	for i := range dp {
		dp[i] = make([]int, nums2len)
	}

	maxCount := 0

	for i := 1; i < nums1len; i++ {
		for j := 1; j < nums2len; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			if dp[i][j] > maxCount {
				maxCount = dp[i][j]
			}
		}
	}

	return maxCount
}

// 解法 2：滚动数组 —— 逐行递推求解法 1 的转移，只保留上一行的状态
// TODO: 实现
// func maxUncrossedLines1(nums1 []int, nums2 []int) int {
// 	return 0
// }
