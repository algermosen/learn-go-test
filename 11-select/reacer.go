package racer

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, defaultTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		winner = url1
		return
	case <-ping(url2):
		winner = url2
		return
	case <-time.After(timeout):
		err = fmt.Errorf("timed out waiting for %s and %s", url1, url2)
		return
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
