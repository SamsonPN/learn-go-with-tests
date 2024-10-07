package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// Fprintf takes in an io.Writer type and writes to it
	fmt.Fprintf(writer, "Hello, %s", name)

	// whereas Printf writes to os.Stdout
	// Printf DEFAULTS to os.Stdout as its io.Writer
	// fmt.Printf("Hello, %s", name)
}

// where else can io.Writer be used?
// for the internet!
// the http.ResponseWriter helps to write your response
// and it implements io.Writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
