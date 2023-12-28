package concurrency

import (
	"testing"
	"reflect"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://www.sverdejot.dev" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	urls := []string{
		"https://www.youtube.com",
		"https://github.com",
		"https://www.sverdejot.dev",
	}

	got := CheckWebsite(mockWebsiteChecker, urls)
	want := map[string]bool{
		"https://www.youtube.com": 	true,
		"https://github.com": 		true,
		"https://www.sverdejot.dev": 	false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := []string{}
	for i := 0; i < 100; i++ {
		urls = append(urls, "url")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}

}
