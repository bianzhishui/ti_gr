package t_039_28_find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// 前缀表 (next数组)
func getNext(next []int, s string) {
	// s = "abababab"
	//  a
	//  0
	//  ab
	//  00
	//  aba
	//  001
	//  abab
	//  0012
	//  abababab    前缀为 [a,ab,aba,abab,ababa,ababab,abababa] 后缀为 [bababab,ababab,babab,abab,bab,ab,b] 所以前缀表最后一个元素为6
	//  00123456
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func GetNext(next []int, s string) {
	getNext(next, s)
}

func StrStr(h, n string) int {
	return strStr(h, n)
}
