package t_024_383_ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	rLen, mLen := len(ransomNote), len(magazine)
	if mLen < rLen {
		return false
	}
	var rList, mList [26]int
	for i := 0; i < mLen; i++ {
		mList[magazine[i]-'a']++
		if i < rLen {
			rList[ransomNote[i]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if rList[i] > mList[i] {
			return false
		}
	}
	return true
}
