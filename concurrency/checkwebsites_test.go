package concurrency

import (
	"reflect"
	"testing"
	"time"
)

const badSite = "https://waaaa.nonsens.grgr"

func MockWebsiteChecker(url string) bool {
	if url == badSite {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	t.Run("should check websites", func(t *testing.T) {
		websites := []string{
			"https://www.google.com",
			"https://www.bing.com",
			badSite,
		}

		expected := make(map[string]bool)
		for _, website := range websites {
			expected[website] = true
		}
		expected[badSite] = false

		actual := CheckWebsites(MockWebsiteChecker, websites)

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("Expected results %v, got results %v", expected, actual)
		}
	})
}

func SlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "bench url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowWebsiteChecker, urls)
	}
}
