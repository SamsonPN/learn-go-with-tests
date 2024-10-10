package concurrency

type WebsiteChecker func(string) bool

// since we don't need names for the fields in the struct
// we can just make them anonymous
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// starts a *goroutine*
		// which is a separate, non-blocking process

		// each iteration of the loop will start a new goroutine
		// concurrent with the current process (the WebsiteChecker function)

		// this causes concurrent map writes
		// if it was just the for-loop, it would have access to the variable url
		// but that variable grabs a new url from the slice of urls

		// but our goroutine only has access to the reference of the url variable (kinda like a pointer)
		// so by the time the goroutine starts, the loop would've probably ended
		// and we would've been on the last url
		// since all goroutines have access to the memory address of the url variable
		// url would've been the same value for all of them

		// go func() {
		// 	results[url] = wc(url)
		// }()

		// so we can instead pass in the url to the goroutine function
		// which would make a copy of url
		// so that we are no longer using just the reference to url

		// but this right here can cause a race condition
		// which is a bug that occurs when the output of our software depends
		// on the timing and sequence of events that we don't have control over
		// we cannot control when each goroutine writes to the results map
		// which causes a fatal error: concurrent map writes to happen
		// you can use go's reace detector to figure this out: `go test -race`

		// go func(u string) {
		// 	results[u] = wc(u)
		// }(url)

		// the race condition can be solved by using a channel
		// a channel is a Go data structure that you can send and receive data
		go func(u string) {
			// this is a `send statement`
			// where we're sending a result to the channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// this ensures that it happens one at a time
	// for writing to the map
	for i := 0; i < len(urls); i++ {
		// this is a receive expression
		// so we're grabbing results from the channel
		// and assigning it
		r := <-resultChannel

		// instead of calling the field name
		// we are calling the types themselves
		results[r.string] = r.bool
	}

	return results
}
