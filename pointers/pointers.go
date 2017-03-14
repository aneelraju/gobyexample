// Go supports pointers, allowing you to pass references to values
// and records
package main

import "fmt"

// arguments passed by value
func zeroval(ival int) {
	ival = 0
}

// arguments passed by reference
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("intial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i syntax gives the memory address of i (pointer to i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// zeroval doesn't change the i in main, but zeroptr does because
	// it has a reference to the memory address for that variable
	fmt.Println("pointer:", &i)
}
