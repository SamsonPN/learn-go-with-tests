package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// starts a *goroutine*
		// which is a separate, non-blocking process

		// each iteration of the loop will start a new goroutine
		// concurrent with the current process (the WebsiteChecker function)

		// this causes concurrent map writes
		// if it was just the for-loop, it would have access to the variable url
		// but that variable grabs a new url from the slice of urls

		// but our goroutine only hass access to the reference of the url variable (kinda like a pointer)
		// so by the time the goroutine starts, the loop would've probably ended
		// and we would've been on the last url
		// since all goroutines have access to the memory address of the url variable
		// url would've been the same value for all of them
		// go func() {
		// 	results[url] = wc(url)
		// }()

		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	// sometimes we would get an empty map
	// b/c the goroutines did not have time to add their result to map
	// for now, we can add this delay
	time.Sleep(2 * time.Second)

	return results
}
