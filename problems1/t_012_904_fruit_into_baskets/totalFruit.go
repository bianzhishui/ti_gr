package t_012_904_fruit_into_baskets

// 滑动窗口
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func totalFruit(fruits []int) int {
	// 收集到的水果树数量
	total := -1
	// 滑动窗口起始位置
	i := 0
	// 篮子 记录遍历时水果类型
	typeMap := map[int]int{}
	for j, fruit := range fruits {
		if _, ok := typeMap[fruit]; ok {
			typeMap[fruit]++
		} else {
			typeMap[fruit] = 1
		}
		// 当记录的类型种类大于 2 时 移动滑动窗口起始位置并将对应记录类型的个数 -1
		for len(typeMap) > 2 {
			typeMap[fruits[i]]--
			if typeMap[fruits[i]] == 0 {
				delete(typeMap, fruits[i])
			}
			i++
		}
		totalTmp := j - (i - 1)
		if total < totalTmp {
			total = totalTmp
		}
	}
	return total
}

// 按块扫描
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func totalFruit1(fruits []int) int {
	// 统计连续出现水果类型的个数
	typeNum := countTypeNumber(fruits)
	// 返回值
	total := -1
	// 遍历
	for i := 0; i < len(typeNum); {
		// 创建集合
		typeMap := map[int]bool{}
		// 返回值临时变量
		totalTmp := 0
		// 标记 break
		flag := true
		// 内层遍历
		for j := i; j < len(typeNum); j++ {
			// 将出现的数字记录到集合中
			typeMap[typeNum[j][0]] = true
			// 临时变量增加
			totalTmp += typeNum[j][1]
			// 若集合长度大于 2 则跳出循环
			if len(typeMap) > 2 {
				flag = false
				i = j - 1
				break
			}
			// 若返回值小于临时变量 则赋值
			if total < totalTmp {
				total = totalTmp
			}
		}
		// 内层遍历结束 说明已经遍历到 typeNum 数组末尾 则直接跳出外层循环
		if flag {
			break
		}
	}
	return total
}

// [3 3 3 1 2 1 1 2 3 3 4]
// [[3 3] [1 1] [2 1] [1 2] [2 1] [3 2] [4 1]]
func countTypeNumber(fruits []int) [][]int {
	var typeNum [][]int
	//for _, num := range fruits {
	//	typeNumLen:=len(typeNum)-1
	//	if typeNumLen==-1 || typeNum[typeNumLen][0] != num {
	//		typeNum = append(typeNum, []int{num, 1})
	//	}else{
	//		typeNum[typeNumLen][1]++
	//	}
	//}
	if len(fruits) == 0 {
		// pass
	} else {
		typeNum = append(typeNum, []int{fruits[0], 1})
		for i := 1; i < len(fruits); i++ {
			typeNumLen := len(typeNum) - 1
			if typeNum[typeNumLen][0] != fruits[i] {
				typeNum = append(typeNum, []int{fruits[i], 1})
			} else {
				typeNum[typeNumLen][1]++
			}
		}
	}
	return typeNum
}
