package t_096_93_restore_ip_addresses

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRTestR(t *testing.T) {
	var (
		s        = "25525511135"
		expected = []string{
			"255.255.11.135",
			"255.255.111.35",
		}
		actual []string
	)
	actual = restoreIpAddresses(s)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
