package t_024_383_ransom_note

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	var (
		ransomNote = "aa"
		magazine   = "aab"
		expected   = true
		actual     bool
	)
	actual = canConstruct(ransomNote, magazine)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
