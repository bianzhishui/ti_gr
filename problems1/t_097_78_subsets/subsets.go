package t_097_78_subsets

func subsets(nums []int) [][]int {
	result := [][]int{}
	ints := []int{}
	var backtrack func(index int)
	backtrack = func(index int) {
		// 没有条件 直接添加ints到result中
		temp := make([]int, len(ints))
		copy(temp, ints)
		result = append(result, temp)
		for i := index; i < len(nums); i++ {
			ints = append(ints, nums[i])
			backtrack(i + 1)
			// 回溯
			ints = ints[:len(ints)-1]
		}
	}
	backtrack(0)
	return result
}
