package t_023_242_valid_anagram

import (
	"reflect"
)

// map[uint8/byte]int{}
func isAnagram(s string, t string) bool {
	// 取 s t 长度
	sLen := len(s)
	tLen := len(t)
	// 长度不相等返回 false
	if sLen != tLen {
		return false
	}
	// 定义 map 保存 uint8/byte : int
	strNumMapS := make(map[uint8]int, sLen)
	strNumMapT := make(map[uint8]int, tLen)
	for i := 0; i < sLen; i++ {
		if _, ok := strNumMapS[s[i]]; ok {
			strNumMapS[s[i]]++
		} else {
			strNumMapS[s[i]] = 1
		}
		if _, ok := strNumMapT[t[i]]; ok {
			strNumMapT[t[i]]++
		} else {
			strNumMapT[t[i]] = 1
		}
	}
	return reflect.DeepEqual(strNumMapS, strNumMapT)
}

// [26]int
func isAnagram2(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen != tLen {
		return false
	}
	var sList, tList [26]int
	for i := 0; i < sLen; i++ {
		sList[s[i]-'a']++
		tList[t[i]-'a']++
	}
	return sList == tList
}

// 排序 s t 排完序判断相等
func isAnagram3(s string, t string) bool {
	return false
}
