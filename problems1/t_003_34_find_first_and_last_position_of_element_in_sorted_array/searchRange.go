package t_003_34_find_first_and_last_position_of_element_in_sorted_array

// 循环遍历
func searchRange(nums []int, target int) []int {
	var list []int
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			list = append(list, i)
			break
		}
	}
	if len(list) == 0 {
		return []int{-1, -1}
	}
	for i := len(nums) - 1; i > -1; i-- {
		if target == nums[i] {
			list = append(list, i)
			break
		}
	}
	return list
}

// 二分查找法
// 需要写两个二分查找法 左查找 右查找
// 左查找返回 -1 则切片中无 target 元素 返回 -1,-1
// 否则进行右查找返回 left,right
func searchRange1(nums []int, target int) []int {
	var left = leftSearch(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	var right = rightSearch(nums, target)
	return []int{left, right}
}

// 左查找
func leftSearch(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	for left <= right {
		var middle = left + (right-left)/2
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			// 如果 middle-1 小于 0 说明 middle 指向了最左边的值 返回 middle
			// 如果 middle-1 大于 0 并且 target!=nums[middle-1] 说明 middle 左边的值不为 target 返回 middle
			if middle-1 == -1 || target != nums[middle-1] {
				return middle
				// middle 左边的值等于 target 则将右指针指向 middle 左边的值
			} else if target == nums[middle-1] {
				right = middle - 1
			}
		}
	}
	// 在 nums 中没有找到 target 返回 -1
	return -1
}

// 右查找
func rightSearch(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	for left <= right {
		var middle = left + (right-left)/2
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			// 如果 middle+1 等于 nums 长度 说明 middle 指向了最右边的值 返回 middle
			// 如果 middle+1 不等于 nums 长度 并且 target!=nums[middle+1] 说明 middle 右边的值不为 target 返回 middle
			if middle+1 == len(nums) || target != nums[middle+1] {
				return middle
				// middle 右边的值等于 target 则将左指针指向 middle 右边的值
			} else if target == nums[middle+1] {
				left = middle + 1
			}
		}
	}
	// 在 nums 中没有找到 target 返回 -1
	return -1
}
