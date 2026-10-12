package t_010_977_squares_of_a_sorted_array

// 双指针逆序组数组
// 时间复杂度：O(n)
func sortedSquares(nums []int) []int {
	// 指向新数组最后一个位置
	var pointer = len(nums) - 1
	// 定义左右指针
	left, right := 0, pointer
	// 定义新数组用来返回
	numsR := make([]int, pointer+1)
	// 遍历
	for left <= right {
		// 计算左右值的平方
		numL, numR := nums[left]*nums[left], nums[right]*nums[right]
		// 如果左值大于右值赋值到新数组的 pointer 位置左指针右移
		if numL > numR {
			numsR[pointer] = numL
			left++
			// 如果右值大于等于左值赋值到新数组的 pointer 位置右指针左移
		} else {
			numsR[pointer] = numR
			right--
		}
		// 新数组下一个位置
		pointer--
	}
	// 返回新数组
	return numsR
}

// 归并排序
func sortedSquares1(nums []int) []int {
	// 定义指针和切片大小
	numJ, size := 0, len(nums)
	// 定义返回切片 起始位置 0 大小 size
	numR := make([]int, 0, size)
	// 遍历寻找正负数交界处 最后 numJ 指向一个正整数
	for ; numJ < size; numJ++ {
		if nums[numJ] >= 0 {
			break
		}
	}
	// numN 指向 numJ 左侧的负数
	var numN = numJ - 1
	// 遍历 四种情况
	for numN >= 0 || numJ < size {
		// 负数指针小于 0 则依次将正整数的平方添加进返回切片中
		if numN < 0 {
			numR = append(numR, nums[numJ]*nums[numJ])
			numJ++
			// 正整数指针等于 size 则一次将负数的平方添加进返回切片中
		} else if numJ == size {
			numR = append(numR, nums[numN]*nums[numN])
			numN--
			// 若负数指针不小于 0 且正整数指针小于 size 则将指针两数平方作比较 将较小的数添加进返回切片中 并移动那个指针
		} else if nums[numJ]*nums[numJ] <= nums[numN]*nums[numN] {
			numR = append(numR, nums[numJ]*nums[numJ])
			numJ++
			// 若负数指针不小于 0 且正整数指针小于 size 则将指针两数平方作比较 将较小的数添加进返回切片中 并移动那个指针
		} else {
			numR = append(numR, nums[numN]*nums[numN])
			numN--
		}
	}
	// 返回切片
	return numR
}
