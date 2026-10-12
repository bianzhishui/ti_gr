package t_005_367_valid_perfect_square

// 循环遍历
func isPerfectSquare(num int) bool {
	for i := 0; i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// 二分查找法
func isPerfectSquare1(num int) bool {
	var left = 0
	var right = num
	for left <= right {
		// 取中间值
		var middle = left + (right-left)/2
		// 找到返回 true
		if middle*middle == num {
			return true
			// 若中间值的平方大于 num 则右指针指向中间值左边的值
		} else if middle*middle > num {
			right = middle - 1
			// 若中间值的平方小于 num 则左指针指向中间值右边的值
		} else {
			left = middle + 1
		}
	}
	// 没有找到返回 false
	return false
}
