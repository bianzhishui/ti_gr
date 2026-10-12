package t_103_51_n_queens

import "strings"

func solveNQueens(n int) [][]string {
	result := [][]string{}
	// 创建棋盘 使用 [][]string 结构创建 方便修改棋盘值 初始化棋盘所有位置值为 .
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	// 记录操作的行数
	row := 0
	// 验证皇后位置是否有效
	var isValid = func(col int) bool {
		// 排除同列存在值为Q的情况
		for i := 0; i < row; i++ {
			if chessboard[i][col] == "Q" {
				return false
			}
		}
		// 排除45°位置存在值为Q的情况
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if chessboard[i][j] == "Q" {
				return false
			}
		}
		// 排除135°位置存在值为Q的情况
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if chessboard[i][j] == "Q" {
				return false
			}
		}
		return true
	}
	var backtrack func()
	backtrack = func() {
		// row==n 遍历到叶子节点说明找到了一组答案 加入到result中
		if row == n {
			// 将棋盘chessboard转换为 []string 结构保存到result中
			temp := make([]string, n)
			for index, value := range chessboard {
				temp[index] = strings.Join(value, "")
			}
			result = append(result, temp)
			return
		}
		// 遍历棋盘一行 找到皇后位置
		for i := 0; i < n; i++ {
			// 验证皇后能否在 [row,i] 该位置
			if isValid(i) {
				chessboard[row][i] = "Q"
				row++
				backtrack()
				// 回溯 先回溯row
				row--
				chessboard[row][i] = "."
			}
		}
	}
	backtrack()
	return result
}
