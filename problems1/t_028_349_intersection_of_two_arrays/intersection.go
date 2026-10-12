package t_028_349_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var ret []int
	numsMap := map[int]bool{}
	for _, num := range nums1 {
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = true
		}
	}
	for _, num := range nums2 {
		if val, ok := numsMap[num]; ok {
			if val {
				ret = append(ret, num)
				numsMap[num] = false
			}
		}
	}
	return ret
}
