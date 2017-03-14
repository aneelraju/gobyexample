// Variables are explicitly declared and used by the compiler to e.g. check
// type-correctness of function calls
package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a string = "initial"
	fmt.Println(a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// infer the of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)
}
