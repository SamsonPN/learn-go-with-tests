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
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)

		// we now have a new dependency on time.Sleep
		// which we need to extract and control
		// it also slows down our tests
		// every time we test something for Countdown, we have to wait 3 seconds
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// in our main function, we are initialize a sleeper variable
	// with the DefaultSleeper struct
	// the DefaultSleeper struct implements the Sleeper interface
	// by implementing the Sleep method which actually calls time.Sleep
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
