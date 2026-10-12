package t_027_1002_find_common_characters

func commonChars(words []string) []string {
	ret := []string{}
	wordsList := make([][]int, 0, len(words))
	for _, word := range words {
		var chars [26]int
		for _, char := range word {
			chars[char-'a']++
		}
		wordsList = append(wordsList, chars[:])
	}
	for i := 0; i < 26; i++ {
		minNum := wordsList[0][i]
		for _, wordsL := range wordsList {
			if minNum > wordsL[i] {
				minNum = wordsL[i]
			}
			if wordsL[i] == 0 {
				break
			}
		}
		for j := 0; j < minNum; j++ {
			ret = append(ret, string(rune('a'+i)))
		}
	}
	return ret
}
