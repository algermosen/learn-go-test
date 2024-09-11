package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	duration1 := measureDuration(url1)
	duration2 := measureDuration(url2)

	if duration1 < duration2 {
		return url1
	}
	return url2
}

func measureDuration(url2 string) time.Duration {
	start2 := time.Now()
	http.Get(url2)
	return time.Since(start2)
}
