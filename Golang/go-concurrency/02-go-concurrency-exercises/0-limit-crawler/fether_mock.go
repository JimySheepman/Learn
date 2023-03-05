package main

import "fmt"

type MockFetcher map[string]*mockResult

type mockResult struct {
	body string
	urls []string
}

func (f MockFetcher) Fetch(url string) (string, []string, error) {
	fetchSignalInstance() <- true
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = MockFetcher{
	"http://golang.org/": &mockResult{
		body: "The Go Programing Language",
		urls: []string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &mockResult{
		body: "Packages",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &mockResult{
		body: "Packages fmt",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &mockResult{
		body: "Packages os",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

var fetchSignal chan bool

func fetchSignalInstance() chan bool {
	if fetchSignal == nil {
		fetchSignal = make(chan bool, 1000)
	}
	return fetchSignal
}
