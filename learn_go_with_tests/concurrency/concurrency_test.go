package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "prefix://something"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"https://quii.gitbook.io/learn-go-with-tests",
		"prefix://something",
	}

	want := map[string]bool{
		"http://google.com":                           true,
		"https://quii.gitbook.io/learn-go-with-tests": true,
		"prefix://something":                          false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
