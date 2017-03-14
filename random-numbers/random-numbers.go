// Go's math/rand package provides pseudorandom number generation
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// returns int n, 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// for other ranges 5.0 <= f < 10/0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// To produce varying sequences, give it a seed that changes
	// This is not safe to use for random numbers you intend to
	// be secret, us crypto/rand for those
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it produces the same
	// sequence of random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

}
