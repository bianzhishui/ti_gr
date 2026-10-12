package t_031_454_4sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ret := 0
	numsMap := map[int]int{}
	for _, i := range nums1 {
		for _, j := range nums2 {
			numsMap[i+j]++
			//k := i+j
			//if _,ok:=numsMap[k];ok{
			//	numsMap[k]++
			//}else{
			//	numsMap[k]=1
			//}
		}
	}
	for _, i := range nums3 {
		for _, j := range nums4 {
			if v, ok := numsMap[-i-j]; ok {
				ret += v
			}
		}
	}
	return ret
}
