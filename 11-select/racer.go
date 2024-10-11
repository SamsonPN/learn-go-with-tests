package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// this allows us to just call the racer
// without having to specify a timeout since we can just use
// the 10 second one by default

// this is moreso for users/the first test
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// doesn't feel good to have to wait 10s for every test
// so we can adjust the timeout duration for tests
// but keep it at 10s for our real code

// this is used moreso for testing purposes
// so that we don't have to wait the full 10+ seconds
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// normally, you can wait for values to be sent to a channel with
	// myVar := <-ch which is a *blocking call*

	// but with select, you can wait on *multiple channels!*
	// first one to send a value wins and its case gets executed

	// in this case, we set up 2 channels for each of the URLs
	// whichever one writes to its channel first will have its case executed
	// and returns the fast URL
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		// time.After is useful when using select
		// there are times when you write code that blocks you forever
		// if a ch never returns a value
		// but time.After returns a ch, like our ping func
		// and it will send a signal after the amount of time you define
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}

func ping(url string) chan struct{} {
	// we don't care what type is sent to the channel
	// we just want a signal that we are done
	// so closing the channel works

	// if we don't care about the type, we can use struct{}
	// reason being, it uses the smallest amount of data vs something like a bool

	// always use make for creating channels
	// if you used var ch chan struct{}, channels zero-value is nil
	// and you can't send to nil channels
	ch := make(chan struct{})

	// this will send a signal into channel once we have completed
	// get
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

// we do not want to actually test our code on real websites
// they could be slow or don't work
// and you can't test edge cases with them
// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
