package t_007_26_remove_duplicates_from_sorted_array

// 双指针法
func removeDuplicates(nums []int) int {
	// 切片大小
	var size = len(nums)
	// 定义指针
	var pointer = 0
	// 遍历
	for i := 0; i < size; i++ {
		// 若指针位置值和遍历位置值相等则继续遍历 找到与指针位置不相等的元素 将此元素放在指针位置的下一位 指针 +1
		if nums[pointer] != nums[i] {
			nums[pointer+1] = nums[i]
			pointer++
		}
	}
	// ## 需要判断 nums 为空的情况 返回 0 否则 返回 pointer+1
	if size == 0 {
		return pointer
	}
	// 返回 pointer+1 pointer 指向新切片最后一个元素 pointer+1 为新切片长度
	return pointer + 1
}
