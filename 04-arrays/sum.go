package sum

func Sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

func SumAll(slices ...[]int) []int {
	sum := make([]int, 0, len(slices))
	for i, s := range slices {
		sum[i] = Sum(s)
	}
	return sum
}
