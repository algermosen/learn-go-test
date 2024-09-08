package iteration

const defaultRepeat = 5

func Repeat(s string, repeat int) (repeated string) {
	if repeat == 0 {
		return ""
	}

	if repeat <= -1 {
		repeat = defaultRepeat
	}

	for i := 0; i < repeat; i++ {
		repeated += s
	}
	return
}
