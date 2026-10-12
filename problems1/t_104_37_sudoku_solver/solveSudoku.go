package t_104_37_sudoku_solver

func solveSudoku(board [][]byte) {
	// 验证[row,col]位置存放c是否合法
	var isValid = func(row, col int, c byte) bool {
		// 排除同行中出现过c的情况
		for i := 0; i < 9; i++ {
			if board[row][i] == c {
				return false
			}
		}
		// 排除同列中出现过c的情况
		for i := 0; i < 9; i++ {
			if board[i][col] == c {
				return false
			}
		}
		// 排除[row,col]位置所在九宫格中出现过c的情况
		startRow, startCol := (row/3)*3, (col/3)*3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == c {
					return false
				}
			}
		}
		return true
	}
	var backtrack func() bool
	backtrack = func() bool {
		// 二维 需要遍历 i行j列中 i*j个元素
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				// 只对 . 元素进行修改
				if board[i][j] == '.' {
					// 从1-9中取数 找到符合题意的数填入 k:='1' k类型为rune
					for k := '1'; k <= '9'; k++ {
						// 判断[i,j]位置填入数字k是否合法
						if isValid(i, j, byte(k)) {
							board[i][j] = byte(k)
							if backtrack() {
								return true
							}
							// 回溯
							board[i][j] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}
	backtrack()
}
