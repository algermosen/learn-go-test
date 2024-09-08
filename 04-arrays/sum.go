package sum

func Sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

func SumAll(slices ...[]int) []int {
	sum := make([]int, len(slices))
	for i, s := range slices {
		sum[i] = Sum(s)
	}
	return sum
}

func SumAllTails(slices ...[]int) []int {
	sum := make([]int, len(slices))
	for i, s := range slices {
		if len(s) == 0 {
			sum[i] = 0
		} else {
			tail := s[1:]
			sum[i] = Sum(tail)
		}
	}
	return sum
}
