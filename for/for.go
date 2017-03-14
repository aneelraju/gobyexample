// For is Go's only looping construct
package main

import "fmt"

func main() {

	i := 1
	// most basic type with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// a classic initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of loop
	// or return from enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// continue to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
