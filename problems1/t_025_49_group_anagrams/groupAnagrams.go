package t_025_49_group_anagrams

func groupAnagrams(strs []string) [][]string {
	// 定义 map 对 strs 字符串列表进行分组
	rMap := map[[26]int][]string{}
	for _, strOne := range strs {
		tmpList := [26]int{}
		for _, char := range strOne {
			tmpList[char-'a']++
		}
		rMap[tmpList] = append(rMap[tmpList], strOne)
	}
	// 定义返回二维列表
	// 将字典存的值添加到返回列表中
	rList := [][]string{}
	for _, sList := range rMap {
		rList = append(rList, sList)
	}
	return rList
}
