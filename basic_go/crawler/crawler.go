package crawler

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func web_crawler() {
	fmt.Println("Exercise: Web Crawler")
	// fetcher is a populated fakeFetcher.
	var fetcher = FakeFetcher{
		"https://golang.org/": &FakeResult{
			Body: "The Go Programming Language",
			Urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &FakeResult{
			Body: "Packages",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &FakeResult{
			Body: "Package fmt",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &FakeResult{
			Body: "Package os",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
	wg.Add(1)
	var crawch = &SafeMapCache{V: make(map[string]bool)}
	go Crawl("https://golang.org/", 4, fetcher, crawch, &wg)
	wg.Wait()

	fmt.Println()
}

type SafeMapCache struct {
	mu sync.Mutex
	V  map[string]bool
}

func (s *SafeMapCache) Get(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.V[key]
	if !ok {
		return false
	}
	return true
}

func (s *SafeMapCache) Set(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.V[key] = true
	return
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FakeFetcher map[string]*FakeResult
type FakeResult struct {
	Body string
	Urls []string
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.Body, res.Urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func Crawl(url string, depth int, fetcher Fetcher, crawch *SafeMapCache, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}
	crawch.Set(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !crawch.Get(u) {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, crawch, wg)
		}
	}
	return
}
