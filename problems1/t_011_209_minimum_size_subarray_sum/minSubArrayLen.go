package t_011_209_minimum_size_subarray_sum

import (
	"math"
	"sort"
)

// 暴力
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func minSubArrayLen(target int, nums []int) int {
	// 数组长度
	l := len(nums)
	// 默认返回值为 数组长度 +1
	ret := l + 1
	for i := range nums {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum >= target {
				ret = j - i + 1
				break
			}
		}
	}
	// 没有找到 返回 0
	if ret == l+1 {
		return 0
	}
	return ret
}

// 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func minSubArrayLen1(target int, nums []int) int {
	// 数组长度
	l := len(nums)
	// 返回值
	ret := l + 1
	// 滑动窗口起始位置
	i := 0
	// 和
	sum := 0
	// j 滑动窗口终止位置
	for j, num := range nums {
		sum += num
		for sum >= target {
			retTmp := j - i + 1
			if retTmp < ret {
				ret = retTmp
			}
			sum -= nums[i]
			i++
		}
	}
	if ret == l+1 {
		return 0
	}
	return ret
}

// 前缀和 + 二分查找
// 时间复杂度: O(n log n)
// 空间复杂度: O(n)
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	// 生成前缀和列表
	// nums=[]int{2, 3, 1, 2, 4, 3}
	// 生成 sums=[]int{0, 2, 5, 6, 8, 12, 15}
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	//fmt.Println("sums:",sums)
	// 遍历 sums
	for i := 1; i <= n; i++ {
		//fmt.Println("i:",i)
		// targetTmp 为 target + i 之前的和
		targetTmp := target + sums[i-1]
		// 二分查找 返回比 targetTmp 小的元素的个数
		//fmt.Println("targetTmp:",targetTmp)
		bound := sort.SearchInts(sums, targetTmp)
		//fmt.Println("bound:",bound)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			// bound-(i-1) 符合条件的数组长度
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
