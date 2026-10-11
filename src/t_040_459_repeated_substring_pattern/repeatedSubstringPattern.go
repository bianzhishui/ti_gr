package t_040_459_repeated_substring_pattern

import (
	t "github.com/bianzhishui/ti_gr/src/t_039_28_find_the_index_of_the_first_occurrence_in_a_string"
)

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
    // abababab
    //  bababa
    // abc
    // abcabc
    //  bcab
	s2 := s + s
	i := t.StrStr(s2[1:2*l-1], s)
	return i != -1
}

func repeatedSubstringPattern1(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	// 前缀表
	next := make([]int, l)
	t.GetNext(next, s)
	// 当字符串s2为字符串s1的重复子串时
	// 例如 s1="abababab" s2="ab" (s1 为本函数的实参)
	// 前缀列表为
	//  a b a b a b a b
	//  0 0 1 2 3 4 5 6 = next
	// next[l-1] == 6 && 8 % (8-6)==0 { true }
	if next[l-1] != 0 && l%(l-next[l-1]) == 0 {
		return true
	}
	return false
}
