package t_090_77_combinations

func combine(n int, k int) [][]int {
	result := [][]int{}
	resultOne := []int{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(resultOne) == k {
			temp := make([]int, k)
			copy(temp, resultOne)
			result = append(result, temp)
			return
		}
		// 剪枝
		for i := index; i <= n-(k-len(resultOne))+1; i++ {
			resultOne = append(resultOne, i)
			backtrack(i + 1)
			resultOne = resultOne[:len(resultOne)-1]
		}
	}
	backtrack(1)
	return result
}
