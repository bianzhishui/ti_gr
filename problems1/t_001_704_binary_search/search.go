package t_001_704_binary_search

// 循环遍历查找
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// 二分查找法
// 方法一
// target 在一个左闭右闭的区间内 [left,right]
// 循环条件为 left<=right
// 若 target==nums[middle] 则已经找到直接返回
// 赋值时 middle-1 middle+1 不包括 middle
func search1(nums []int, target int) int {
	// 定义左右指针
	var (
		left   = 0
		right  = len(nums) - 1
		middle = -1
	)
	// 循环
	for left <= right {
		// 中间值
		middle = left + (right-left)/2
		// 查找值小于中间值 右指针指向中间值左边的值
		if target < nums[middle] {
			right = middle - 1
			// 查找值大于中间值 左指针指向中间值右边的值
		} else if target > nums[middle] {
			left = middle + 1
			// 查找值等于中间值 返回
		} else {
			break
		}
	}
	// 循环结束 没有找到返回 -1
	return middle
}

// 二分查找法
// 方法二
// target 在一个左闭右开的区间内 [left,right)
// 循环条件为 left<right
// 若 target==nums[middle] 则已经找到直接返回
// 若 target<nums[middle] right=middle // 因为左闭右开右指针需要指向中间值
func search2(nums []int, target int) int {
	// 定义左右指针
	var (
		left   = 0
		right  = len(nums)
		middle = -1
	)
	// 循环
	for left < right {
		// 中间值
		middle = left + (right-left)/2
		// 查找值小于中间值 右指针指向中间值左边的值
		if target < nums[middle] {
			right = middle
			// 查找值大于中间值 左指针指向中间值右边的值
		} else if target > nums[middle] {
			left = middle + 1
			// 查找值等于中间值 返回
		} else {
			break
		}
	}
	// 循环结束 没有找到返回 -1
	return middle
}

// 二分查找法使用条件：
// 1.有序
// 2.无重复值
