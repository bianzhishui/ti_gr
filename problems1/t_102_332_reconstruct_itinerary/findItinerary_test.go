package t_102_332_reconstruct_itinerary

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	var (
		tickets = [][]string{
			{"MUC", "LHR"},
			{"JFK", "MUC"},
			{"SFO", "SJC"},
			{"LHR", "SFO"},
		}
		expected = []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
		actual   []string
	)
	actual = findItinerary(tickets)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
