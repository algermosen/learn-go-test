package iteration

const repeatCount = 5

func Repeat(s string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}
	return
}
