package concurrency

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !maps.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100) // <- 100 is an arbitrary number that we previously decided it's enough

	// The urls to be tested are not important, the importante behaviour is the time spent in each call
	// therefore, we will fill the urls array with an arbitrary url
	for i := 0; i < len(urls); i++ {
		urls[i] = "an url" // <- Proof that the url does not matter. Remember you're testing with slowStubWebsiteChecker
	}

	b.ResetTimer() // Reset the time to just take the function execution and not the preparation step

	// To make a benchmark test, I need to loop the function execution with the b.N
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls) // <- Since there are urls needed, we have to declare them
	}
}
