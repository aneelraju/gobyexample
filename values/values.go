// Go has various value types including string, integers, floats, booleans etc
package main

import "fmt"

func main() {
	// string can be added together with +
	fmt.Println("go" + "lang")

	// integers
	fmt.Println("1+1 =", 1+1)
	//floats
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
