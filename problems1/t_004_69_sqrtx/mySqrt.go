package t_004_69_sqrtx

// 循环遍历
func mySqrt(x int) int {
	for i := 1; true; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
	return -1
}

// 二分查找法
func mySqrt1(x int) int {
	var left = 1
	var right = x
	for left <= right {
		// 中间值
		var middle = left + (right-left)/2
		// 若正好中间值的平方等于 x 返回 middle
		if middle*middle == x {
			return middle
			// 若中间值的平方大于 x
		} else if middle*middle > x {
			// 首先判断中间值左边值的平方 是否小于 x 若小于 x 返回 middle-1 否则 右指针指向中间值左边的值
			if (middle-1)*(middle-1) < x {
				return middle - 1
			}
			right = middle - 1
			// 若中间值的平方小于 x
		} else {
			// 首先判断中间值右边值的平方 是否大于 x 若大于 x 返回 middle 否则 左指针指向中间值右边的值
			if (middle+1)*(middle+1) > x {
				return middle
			}
			left = middle + 1
		}
	}
	// x 是非负整数
	// 正整数必有平方根
	// 若程序运行到此 且 x 为正整数 则 x=0 0 的平方根为 0 返回 0
	return 0
}
