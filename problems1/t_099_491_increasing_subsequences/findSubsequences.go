package t_099_491_increasing_subsequences

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	ints := []int{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(ints) > 1 {
			temp := make([]int, len(ints))
			copy(temp, ints)
			result = append(result, temp)
		}
		// 去重 把map当作set 或 使用数组 题意中nums元素有范围可以使用数组
		uset := map[int]struct{}{}
		// 同一层去重 (回溯之后去重)
		for i := index; i < len(nums); i++ {
			// if _, ok := uset[nums[i]]; len(ints) != 0 && nums[i] < ints[len(ints)-1] || ok {
			// 	continue
			// }
			if _, ok := uset[nums[i]]; (len(ints) == 0 || nums[i] >= ints[len(ints)-1]) && !ok {
				uset[nums[i]] = struct{}{}
				ints = append(ints, nums[i])
				backtrack(i + 1)
				ints = ints[:len(ints)-1]
			}
		}
	}
	backtrack(0)
	return result
}
