package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// first we define our dependency as an interface
type Sleeper interface {
	Sleep()
}

// spies are a kind of *mock* that can record how a dependency is used
// in this case, we are seeing if time.Sleep is being called 3 times

// this slice of strings helps us determine the ORDER of operations
// we want write, sleep, write, sleep, etc
// our original test with just spySleeper.calls only
// tests if we sleep 3 times, not if we did it in the correct order
type SpyCountdownOperations struct {
	Calls []string
}

// makes SpyCountdownOperations implement the Sleeper interface
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// makes SpyCountdownOperations implement the io.Writer interface
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

// now we want our sleeper to be able to set its own duration
// and sleep
// duration = time slept
// sleep = way to pass in a sleep function
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

// SpyTime now implements sleeper
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)

		// we used to have a dependency on time.Sleep
		// which we needed to extract and control
		// and slowed down our tests
		// every time we test something for Countdown, we have to wait 3 seconds

		// so by creating the Sleeper interface and having methods for sleep
		// we can simulate it by just calling sleep 3 times without having to wait
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// allows us to pass in our own sleep function
	// which allows us to test if it's actually sleeping like we want it to
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
