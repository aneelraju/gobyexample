// Go's structs are typed collections of fields
// They are useful for grouping data together to form
// records
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	// syntax to create new struct
	fmt.Println(person{"Bob", 20})

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// an & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// access struct with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can also use dots with struct pointers - the pointers
	// are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

}
