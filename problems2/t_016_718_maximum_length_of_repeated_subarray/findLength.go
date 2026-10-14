package t_016_718_maximum_length_of_repeated_subarray

import "slices"

// 解法 1：二维动态规划 —— dp[i][j] 表示以 nums1[i-1]、nums2[j-1] 结尾的最长公共子数组长度
// TODO: 实现
func findLength(nums1 []int, nums2 []int) int {

	nums1 = slices.Insert(nums1, 0, 0)
	nums2 = slices.Insert(nums2, 0, 0)

	nums1len := len(nums1)
	nums2len := len(nums2)

	dp := make([][]int, nums1len)
	result := 0

	for i := range dp {
		dp[i] = make([]int, nums2len)
	}

	for i := 1; i < nums1len; i++ {
		for j := 1; j < nums2len; j++ {

			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if dp[i][j] > result {
				result = dp[i][j]
			}

		}
	}

	return result

}

// 解法 2：滚动数组 —— 只保留前一行的状态，内层倒序遍历，避免覆盖上一行尚未使用的值
// TODO: 实现
// func findLength1(nums1 []int, nums2 []int) int {
// 	return 0
// }
