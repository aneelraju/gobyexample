// Environment variables are a universal mechanism for conveying
// configuration information to Unix programs
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use os.Setenv
	// To get a value use os.Getenv
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use os.Environ to list all key/values pairs in the environment
	// This returns a slice of strings in the form KEY=value. You
	// can strings.Split them to get the key and value
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
