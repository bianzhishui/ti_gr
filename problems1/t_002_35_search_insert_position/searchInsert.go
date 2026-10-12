package t_002_35_search_insert_position

// 循环遍历
func searchInsert(nums []int, target int) int {
	var i = 0
	for i < len(nums) {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			return i
		} else {
			i++
		}
	}
	return i
}

// 二分查找法
// 方法一
// 若 target==nums[middle] 则直接返回 middle
// 循环结束 返回 right+1 // 因为在最后一次循环开始时 left==right 指向同一个值 经过判断以后 right 会指向比 target 小的值 所以 返回 right+1
//
// 四种情况:
// 1.目标值在数组所有元素之前 [0,-1]
// 2.目标值等于数组中某一个元素 return middle;
// 3.目标值插入数组中的位置 [left,right],return right+1
// 4.目标值在数组所有元素之后的情况 [left,right],return right+1
func searchInsert1(nums []int, target int) int {
	// 定义左右指针
	var left = 0
	var right = len(nums) - 1
	// 循环
	for left <= right {
		// 中间值
		middle := left + (right-left)/2
		// 查找值小于中间值 右指针指向中间值左边的值
		if target < nums[middle] {
			right = middle - 1
			// 查找值大于中间值 左指针指向中间值右边的值
		} else if target > nums[middle] {
			left = middle + 1
			// 查找值等于中间值 返回
		} else {
			return middle
		}
	}
	// 循环结束 没有找到返回 right+1
	return right + 1
}

// 二分查找法
// 方法二
// 若 target==nums[middle] 则直接返回 middle
// 循环结束 left==right 返回 right 或 left // 因为在最后一次循环开始时 left+1==right 指向相邻的两个值 midele 指向 left 经过判断以后 left==right 指向比 target 大的值 所以 返回 right 或 left
//
// 四种情况:
// 1.目标值在数组所有元素之前 [0,0)
// 2.目标值等于数组中某一个元素 return middle;
// 3.目标值插入数组中的位置 [left,right),return right
// 4.目标值在数组所有元素之后的情况 [left,right),return right
func searchInsert2(nums []int, target int) int {
	// 定义左右指针
	var left = 0
	var right = len(nums)
	// 循环
	for left < right {
		// 中间值
		middle := left + (right-left)/2
		// 查找值小于中间值 右指针指向中间值左边的值
		if target < nums[middle] {
			right = middle
			// 查找值大于中间值 左指针指向中间值右边的值
		} else if target > nums[middle] {
			left = middle + 1
			// 查找值等于中间值 返回
		} else {
			return middle
		}
	}
	// 循环结束 没有找到返回 right 或 left
	return right
}
