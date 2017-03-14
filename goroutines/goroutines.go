// A goroutine is a lightweight thread of execution
package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// you can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("golang")

	// The Scanln code requires we press a key before the program exits
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
