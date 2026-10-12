package t_009_283_move_zeroes

func moveZeroes(nums []int) {
	var pointer = 0
	for _, v := range nums {
		if v != 0 {
			nums[pointer] = v
			pointer++
		}
	}
	for ; pointer < len(nums); pointer++ {
		nums[pointer] = 0
	}
}

// 双指针
func moveZeroes1(nums []int) {
	// 定义两个指针
	// 左指针左侧均为非零数
	// 右指针左侧直到左指针处均为零
	left, right := 0, 0
	// 遍历
	for right < len(nums) {
		// 右指针找到非零数与左指针交换 左指针+1
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		// 右指针找到数为零 右指针+1
		right++
	}
}
