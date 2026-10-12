package t_029_202_happy_number

func isHappy(n int) bool {
	numMap := map[int]struct{}{n: {}}
	sumNum := func(num int) (sum int) {
		for ; num != 0; num = num / 10 {
			sum += (num % 10) * (num % 10)
		}
		return
	}
	for {
		sum := sumNum(n)
		if sum == 1 {
			return true
		}
		if _, ok := numMap[sum]; ok {
			return false
		} else {
			numMap[sum] = struct{}{}
		}
		n = sum
	}
}
