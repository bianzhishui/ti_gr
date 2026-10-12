package t_015_54_spiral_matrix

// 模拟
// 时间复杂度: O(mn)
// 空间复杂度: O(mn)
func spiralOrder(matrix [][]int) []int {
	// 若长度等于 0 返回空数组
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 取二维数组行数和列数
	rows, columns := len(matrix), len(matrix[0])
	// 创建和参数相同的二维数组用来标记位置是否读取过
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}
	var (
		// 数组大小
		total = rows * columns
		// 定义返回的一维数组
		order = make([]int, total)
		// 标记数组位置
		row, column = 0, 0
		// 定义方向数组
		directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		// 取方向数组位置
		directionIndex = 0
	)
	// 遍历
	for i := 0; i < total; i++ {
		// 给一维数组赋值
		order[i] = matrix[row][column]
		// 赋值完成记录位置已经读取过
		visited[row][column] = true
		// 取下一个位置用来判断是否有效
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		// 若行数或列数大于最大小于最小 或 当前位置是读取过的 则 取方向数组的下一位置
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		// 改变位置
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

// 模拟
// 时间复杂度: O(mn)
// 空间复杂度: O(1)
func spiralOrder1(matrix [][]int) []int {
	// 若数组为空返回空数组
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		// 行数和列数
		rows, columns = len(matrix), len(matrix[0])
		// 一维数组用来返回
		order = make([]int, rows*columns)
		// 标记一维数组的位置
		index = 0
		// 定义边界 最左 最右 最上 最下
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	// 遍历
	for left <= right && top <= bottom {
		// 先取最左的位置向右遍历
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		// 遍历到最右边 然后向下遍历
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		// 判断数值有效性 然后向左遍历 再向上遍历
		if left < right && top < bottom {
			// 向左遍历
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			// 向上遍历
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		// 遍历完一层 最左值和最上值++ 最右值和最下值--
		left++
		right--
		top++
		bottom--
	}
	return order
}
