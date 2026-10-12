package t_092_17_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	// 数字到字母的映射
	letterMap := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	bytes := []byte{}
	k := len(digits)
	// backtrack index为digits切片中需要处理的字符下标
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(bytes) == k {
			result = append(result, string(bytes))
			return
		}
		// 找到需要处理的字符组合
		letter := letterMap[digits[index]-'0']
		for i := 0; i < len(letter); i++ {
			bytes = append(bytes, letter[i])
			backtrack(index + 1)
			bytes = bytes[:len(bytes)-1]
		}

	}
	backtrack(0)
	return result
}

func letterCombinations1(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	letterMap := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	str := ""
	k := len(digits)
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(str) == k {
			result = append(result, str)
			return
		}
		letter := letterMap[digits[index]-'0']
		for i := 0; i < len(letter); i++ {
			str += string(letter[i])
			backtrack(index + 1)
			str = str[:len(str)-1]
		}

	}
	backtrack(0)
	return result
}
