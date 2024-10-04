package main

import "testing"

func TestHello(t *testing.T) {
	// each t.Run is a subtest that corresponds to different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Samson", "")
		want := "Hello, Samson"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Samantha", "French")
		want := "Bonjour, Samantha"
		assertCorrectMessage(t, got, want)
	})
}

// we moved the if got != want block into this function
// this reduces duplicate code since some tests use this assertion

// we pass in t *testing.T to tell the test code to fail

// for these helper functions, we should accept a testing.TB
// b/c it is an interface that *testing.T and *testing.B satisfy
// THIS ALLOWS YOU TO CALL THESE HELPER FUNCTIONS FROM A TEST OR A BENCHMARK
func assertCorrectMessage(t testing.TB, got, want string) {
	// this tests the test suite that this method is a helper
	// this is so that the line number for the error will be in the function call
	// NOT in the test helper itself!!!

	// without this, it would show the fail at line 33
	// rather than at line 10/15 where the test calls this helper function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// if you can't run this, then run
// go mod init hello
