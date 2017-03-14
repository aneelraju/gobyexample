// Defer is used to ensure that a function call is performed later in a program's
// execution, usually for purpose of cleanup
package main

import (
	"fmt"
	"os"
)

func main() {

	// Immediately after getting a file object with createFile, we defer
	// the closing of that file with closeFile. This will be executed at
	// the end of the enclosing function(main), after writeFile has finished
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
