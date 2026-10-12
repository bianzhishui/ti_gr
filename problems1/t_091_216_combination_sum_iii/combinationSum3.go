package t_091_216_combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	resultOne := []int{}
	sum := 0
	var backtrack func(index int)
	backtrack = func(index int) {
		// 剪枝
		if sum > n {
			return
		}
		if len(resultOne) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, resultOne)
				result = append(result, temp)
			}
			return
		}
		// 剪枝: i <= 9-(k-len(resultOne))+1  //若剩余个数不足组成k个元素则剪枝
		for i := index; i <= 9-(k-len(resultOne))+1; i++ {
			sum += i
			resultOne = append(resultOne, i)
			backtrack(i + 1)
			// 回溯
			resultOne = resultOne[:len(resultOne)-1]
			sum -= i
		}
	}
	backtrack(1)
	return result
}
