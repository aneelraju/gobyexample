// A panic typically means something went unexpectedly wrong
// Mostly we use it to fail fast on errors that shouldn't occur
// during normal operation, or that we aren't prepared to handle
// gracefully

package main

import "os"

func main() {

	panic("a problem")

	// A common use of panic is to abort if a function returns an error value
	// that we don't know how to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	// Running the program will cause it to panic, print an error message
	// and goroutine traces, and exit with a non-zero status

	// Note that unlike some languages which use exceptions for handling of
	// many errors, in Go it is idiomatic to use error indication return
	// values wherever possible
}
