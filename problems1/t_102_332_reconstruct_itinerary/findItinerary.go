package t_102_332_reconstruct_itinerary

import "sort"

type to struct {
	place string
	used  bool
}

type tos []*to

func (t tos) Len() int {
	return len(t)
}

func (t tos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tos) Less(i, j int) bool {
	return t[i].place < t[j].place
}

// tickets = {"from", "to"},
func findItinerary(tickets [][]string) []string {
	result := []string{}
	// 航线信息映射 map[from]=[]{{to,false},{...}}
	froms := make(map[string]tos)
	for _, ticket := range tickets {
		if _, ok := froms[ticket[0]]; !ok {
			froms[ticket[0]] = make(tos, 0, 1)
		}
		froms[ticket[0]] = append(froms[ticket[0]], &to{
			place: ticket[1],
			used:  false,
		})
	}
	// 将tos进行排序 因为题意需要返回较小值
	for key := range froms {
		sort.Sort(froms[key])
	}
	result = append(result, "JFK")
	var backtrack func() bool
	backtrack = func() bool {
		// 当tickets长度等于1 时说明有两个机场 也就是机场数等于tickets长度加1 当result等于这个值时说明找到了符合题意的机场顺序
		if len(tickets)+1 == len(result) {
			return true
		}
		// 找下一趟航班
		for _, value := range froms[result[len(result)-1]] {
			// 若place值没有使用过 说明找到了航班
			if value.used == false {
				result = append(result, value.place)
				value.used = true
				if backtrack() {
					return true
				}
				// 回溯
				result = result[:len(result)-1]
				value.used = false
			}
		}
		return false
	}
	backtrack()
	return result
}
