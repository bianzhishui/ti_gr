package t_014_59_spiral_matrix_ii

// 模拟
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func generateMatrix(n int) [][]int {
	// 创建二维数组(矩阵)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 定义方向数组
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 定义当前位置
	x, y := 0, 0
	// 当前方向为初始方向
	z := 0
	// 取方向
	direction := directions[z]
	// 生成数组
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		// 定义临时变量指向该方向的下一位置
		xTmp, yTmp := x+direction[0], y+direction[1]
		// 判断该位置有效性
		if xTmp < 0 || xTmp >= n || yTmp < 0 || yTmp >= n || matrix[xTmp][yTmp] > 0 {
			// 改变方向时利用取余 因为生成矩阵会多次改变方向
			z = (z + 1) % 4
			direction = directions[z]
		}
		x += direction[0]
		y += direction[1]
	}
	return matrix
}

// 按层模拟
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func generateMatrix1(n int) [][]int {
	// top 最顶层 bottom 最底层
	top, bottom := 0, n-1
	// left 最左 right 最右
	left, right := 0, n-1
	// 元素个数
	num := 1
	// 元素总个数
	tar := n * n
	// 返回的二维数组
	matrix := make([][]int, n)
	// 初始化二维数组
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 循环赋值 num 记录赋值个数 若大于总个数赋值则赋值完
	for num <= tar {
		// 先赋值最顶层 从左向右赋值 num 记录个数
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		// 顶层赋值完 top++
		top++
		// 赋值最右层
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 赋值最底层
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// 赋值最左层
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
