package t_006_27_remove_element

// 两层循环 第一层循环遍历 第二层循环更新数组
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func removeElement(nums []int, val int) int {
	// 切片大小
	var size = len(nums)
	// 遍历
	for i := 0; i < size; i++ {
		// 若 nums[i]==val
		if nums[i] == val {
			// 将第 i 个元素移动到切片最后
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			// 因为经过移动 i 位置上的数为原来 i 后面的数 该数没有经过 nums[i]==val 判断 所以 i--
			i--
			// 切片大小 -1
			size--
		}
	}
	// 遍历完返回 size
	return size
}

// 双指针法
func removeElement1(nums []int, val int) int {
	// 定义左右指针
	left, right := 0, len(nums)-1
	// 循环条件 left<=right
	for left <= right {
		// 若左指针指向的值等于 val 则将右指针值给左指针位置 右指针左移
		if nums[left] == val {
			nums[left] = nums[right]
			right--
			// 若左指针指向的值不等于 val 则左指针右移
		} else {
			left++
		}
		// ## 不对右指针的值进行判断 在循环中处理
	}
	// 返回 left
	return left
}

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElement2(nums []int, val int) int {
	// 定义指针
	var pointer = 0
	// 遍历
	for _, v := range nums {
		// 若 v 不等于 val 则赋值给 pointer 指针位置上并 pointer++
		if v != val {
			nums[pointer] = v
			pointer++
		}
	}
	// 返回 pointer pointer 也为新切片大小
	return pointer
}
