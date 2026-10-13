package t_014_53_maximum_subarray

// 解法 1：线性动态规划（Kadane）—— 以当前位置结尾的最大子数组和
// TODO: 实现
func maxSubArray(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	dp := nums[0]
	result := nums[0]

	for i := 1; i < numsLen; i++ {

		if dp+nums[i] > nums[i] {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
		}

		if dp > result {
			result = dp
		}

	}

	return result

}

// 解法 2：分治法 —— 跨越中点的最大子数组与左右两半递归结果取最大
// TODO: 实现
// func maxSubArray1(nums []int) int {
// 	return 0
// }
