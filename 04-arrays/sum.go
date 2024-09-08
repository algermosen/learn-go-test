package sum

func Sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
