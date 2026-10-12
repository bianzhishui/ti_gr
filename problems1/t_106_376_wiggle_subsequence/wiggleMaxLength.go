package t_106_376_wiggle_subsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		up   = 1
		down = 1
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(up, down)
}

func wiggleMaxLength2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		up   []int = make([]int, len(nums))
		down []int = make([]int, len(nums))
	)
	up[0] = 1
	down[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i] = up[i-1] + 1
			up[i] = up[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[len(nums)-1], down[len(nums)-1])
}
