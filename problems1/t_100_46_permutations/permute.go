package t_100_46_permutations

func permute(nums []int) [][]int {
	result := [][]int{}
	ints := []int{}
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		// 找到符合题意的组合
		if len(ints) == len(nums) {
			temp := make([]int, len(ints))
			copy(temp, ints)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 说明使用过 nums[i] 元素
			if used[i] {
				continue
			}
			ints = append(ints, nums[i])
			used[i] = true
			backtrack()
			// 回溯
			ints = ints[:len(ints)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}
