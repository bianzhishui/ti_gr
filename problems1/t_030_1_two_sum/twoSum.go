package t_030_1_two_sum

func twoSum(nums []int, target int) []int {
	var ret []int
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		j := target - num
		if k, ok := numsMap[j]; ok {
			ret = []int{k, i}
			break
		}
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = i
		}
	}
	return ret
}
