// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays
package main

import "fmt"

func main() {

	// unlike arrays, slices are typed only by the elements they contain
	// to create an empty slice with non-zero length, use the builtin make
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// we can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// slices support several more operations than arrays
	// like builtin append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// syntax slice[low:high]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// we can delcare and initialize a variable for slice in a single
	// line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices can be compused into multi-dimensional data structures
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
