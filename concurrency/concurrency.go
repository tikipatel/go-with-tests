package concurrency

// WebsiteChecker is a type
type WebsiteChecker func(string) bool

// CheckWebsites is a function that...
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
